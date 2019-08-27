<template>
 <div class="app-container">
    <div class="filter-container">
        <div class="filter-item">
            <el-date-picker v-model="listQuery.startEnd" type="datetimerange" range-separator="至" start-placeholder="开始时间" end-placeholder="结束时间">
            </el-date-picker>
        </div>
        <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
    </div>
    <el-table
            v-loading="listLoading"
            :key="tableKey"
            :data="logData"
            border
            fit
            highlight-current-row
            style="width: 100%;margin-top:15px;">
            <el-table-column label="日志日期" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.beginTime | parseDateFilter }}</span>
                </template>
            </el-table-column>
            <el-table-column label="日志类型" align="center" width="200px">
                <template slot-scope="scope">
                    <span>{{ scope.row.statType}}</span>
                </template>
            </el-table-column>
            <el-table-column label="日志数量" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.statCount}}</span>
                </template>
            </el-table-column>
            <el-table-column label="最后更新时间" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.updateTime | parseTimeFilter}}</span>
                </template>
            </el-table-column>
        </el-table>

        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="totalCount" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>
 </div>    
    
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import { getPlayerStats } from "@/api/log";
import { parseTime } from "@/utils/index";
export default {
  name: "PlayerStats",
  directives: {
    waves
  },
  filters: {
    parseDateFilter: function(value) {
      return parseTime(value, "{y}-{m}-{d}");
    },
    parseTimeFilter: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
    }
  },
  created() {
    this.initMetaData();
    // this.getList();
  },
  data() {
    return {
      tableKey: 1,
      listQuery: {
        beginTime: undefined,
        endTime: undefined,
        pageIndex: 1,
        startEnd:[]
      },
      listLoading: false
    };
  },
  methods: {
    handleFilter(e) {
      this.listQuery.pageIndex = 1;
      if (this.listQuery.startEnd && this.listQuery.startEnd.length == 2) {
        this.listQuery.beginTime = this.listQuery.startEnd[0].valueOf();
        this.listQuery.endTime = this.listQuery.startEnd[1].valueOf();
      }
      this.loadData();
    },
    handleCurrentChange(e) {
      this.listQuery.pageIndex = e;
      this.loadData();
    },
    initMetaData() {
      let startDate = new Date();
      startDate = new Date(
        startDate.getFullYear(),
        startDate.getMonth(),
        startDate.getDate()
      );
      let endDate = new Date();
      endDate.setDate(endDate.getDate() + 1);
      endDate = new Date(
        endDate.getFullYear(),
        endDate.getMonth(),
        endDate.getDate()
      );
      this.listQuery.startEnd = [startDate, endDate];
    },
    loadData() {
      this.listLoading = true;
      getPlayerStats(this.listQuery).then(res => {
        this.logData = res.itemArray;
        this.totalCount = res.totalCount;
        this.listLoading = false;
      });
    }
  }
};
</script>

