<template>
 <div class="app-container">
    <div class="filter-container">
        <el-select v-model="listQuery.tableName" placeholder="日志类型" style="width: 120px" class="filter-item" @change="handleLogTypeChange">
            <el-option v-for="item in metaLogType" :key="item.key" :label="item.value" :value="item.key"/>
        </el-select>
        <el-select v-model="listQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handlePlatformChange">
            <el-option v-for="item in platformList" :key="item.centerPlatformId" :label="item.centerPlatformName" :value="item.centerPlatformId" />
        </el-select>
        <el-select v-model="listQuery.sid" collapse-tags placeholder="服务器" clearable style="width: 180px" class="filter-item" @change="handleServerChange">
            <el-option v-for="item in serverList" :key="item.id" :label="item.serverName" :value="item.id"/>
        </el-select>
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
            <el-table-column label="日志时间" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.logTime | parseTimeFilter }}</span>
                </template>
            </el-table-column>
            <el-table-column label="平台ID" width="80px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.platform}}</span>
                </template>
            </el-table-column>
            <el-table-column label="服务器ID" width="100px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.serverId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="服务器类型" width="100px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.serverType}}</span>
                </template>
            </el-table-column>
            <el-table-column v-for="(item,index) in metaColumnArray" :key="index" :label="item.lab" v-if="index < metaColumnArray.length-1" width="100px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row[item.key] | commonFilter(item.showType) }}</span>
                </template>
            </el-table-column>
            <el-table-column v-for="(item,index) in metaColumnArray" :key="index"  :label="item.lab" v-if="index == metaColumnArray.length-1" min-width="100px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row[item.key] | commonFilter(item.showType) }}</span>
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
import { getCenterPlatList, getAllCenterServerList } from "@/api/center";
import { getLog, getLogMeta, getLogMetaMsgList } from "@/api/log";
import { parseTime } from "@/utils/index";
export default {
  name: "ServerLogList",
  directives: {
    waves
  },
  filters: {
    parseTimeFilter: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
    },
    commonFilter: function(value, type) {
      if (type == "datetime") {
        return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
      }
      if (type == "byte") {
        return binaryToStr(value);
      }
      if (type == "normal") {
        return value;
      }
      return value;
    }
  },
  created() {
    this.initMetaData();
    // this.getList();
  },
  data() {
    return {
      tableKey: 1,
      //基础元数据
      metaLogType: [], //日志列表
      platformList: [], //平台列表
      serverList: [], //服务器列表
      listQuery: {
        tableName: undefined,
        beginTime: undefined,
        endTime: undefined,
        platformId: undefined,
        sid: undefined,
        serverType: -1,
        startEnd: [],
        serverId: undefined,
        pageIndex: 1,
        playerId: undefined
      },
      logData: [],
      totalCount: 0,
      listLoading: false,
      tableMetaMap: new Map(),
      metaColumnArray: []
    };
  },
  methods: {
    handleLogTypeChange(e) {
      this.loadMetaColumn(e);
    },
    handlePlatformChange(e) {
      this.listQuery.sid = undefined;
      this.listQuery.serverType = -1;
      this.listQuery.serverid = undefined;
      getAllCenterServerList(e).then(res => {
        this.serverList = res.itemArray;
      });
      //   console.log(this.listQuery);
    },
    handleServerChange(e) {
      if (!e) {
        this.listQuery.serverType = -1;
        this.listQuery.serverId = undefined;
        return;
      }
      let item = this.findServerItem(e);
      if (item) {
        this.listQuery.serverId = item.serverId;
        this.listQuery.serverType = item.serverType;
      }
      //   console.log(this.listQuery);
    },
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
      getLogMetaMsgList(0).then(res => {
        this.metaLogType = res;
        if (this.metaLogType && this.metaLogType.length > 0) {
          this.listQuery.tableName = this.metaLogType[0].key;
          this.loadMetaColumn(this.metaLogType[0].key);
        }
      });

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
      getCenterPlatList().then(res => {
        this.platformList = res.itemArray;
      });
    },
    loadMetaColumn(e) {
      let metaData = this.tableMetaMap.get(e);
      if (metaData) {
        this.metaColumnArray = metaData;
        return;
      }
      getLogMeta(e, 0).then(res => {
        this.tableMetaMap.set(e, res);
        this.metaColumnArray = res;
      });
    },
    loadData() {
      if (!this.listQuery.tableName) {
        this.$message.error("日志类型不能为空");
        return;
      }
      this.listLoading = true;
      getLog(this.listQuery).then(res => {
        this.logData = res.itemArray;
        this.totalCount = res.totalCount;
        this.listLoading = false;
      });
    },
    findServerItem(serverid) {
      const server = this.serverList.find(n => {
        return n.id == serverid;
      });
      if (server) {
        return server;
      }
      return undefined;
    }
  }
};
</script>

