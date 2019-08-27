<template>
  <div>
    <div class="filter-container">
      <el-select v-model="listQuery.platformId" class="filter-item" clearable placeholder="中心平台">
         <el-option v-for="citem in centerPlatformList" :key="citem.centerPlatformId" :label="citem.centerPlatformName" :value="citem.centerPlatformId"/>
      </el-select>
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
    </div>

    <el-table
      v-loading="listLoading"
      :key="tableKey"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;margin-top:15px;">
      <el-table-column label="日志ID" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column label="中心平台ID" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ getPlatformName(scope.row.centerPlatformId) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="中心服务序号" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.centerServerId }}</span>
        </template>
      </el-table-column>
       <el-table-column label="成功标志" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.successFlag | parseYesOrNo }}</span>
        </template>
      </el-table-column>
       <el-table-column label="发送类型" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.kindType | parseMarryKindType }}</span>
        </template>
      </el-table-column>
       <el-table-column label="异常信息" align="center" min-width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.failMsg }}</span>
        </template>
      </el-table-column>
       <el-table-column label="发送时间" align="center" width="160px">
        <template slot-scope="scope">
          <span>{{ scope.row.updateTime | parseTimeFilter }}</span>
        </template>
      </el-table-column>

      <el-table-column fixed="right" label="操作" align="center" width="140" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button v-if="scope.row.successFlag == 0" type="primary" size="mini" @click="handleSend(scope.row)">发送</el-button>
        </template>
      </el-table-column>
    </el-table>

    <div class="pagination-container" style="margin-top:15px;">
      <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"/>
    </div>

    <el-dialog :visible.sync="dialogPvVisible" title="是否确认发送" width="30%">
      <div>
        是否确认发送?
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogPvVisible = false">取消</el-button>
        <el-button type="primary" @click="sendData">发送</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import { 
    getCenterPlatformMarrySetLogList,
    updateCenterPlatformMarrySetSend 
} from "@/api/centerPlatform";
import {
  serverTypeList,
  serverTagList,
  serverStatusList
} from "@/types/center";
import { parseTime } from "@/utils/index";
import { getCenterPlatList } from "@/api/center";

export default {
  name: "CenterPlatformLogList",
  directives: {
    waves
  },
  filters: {
    parseTimeFilter: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
    },
    parseYesOrNo: function(value) {
      if (value) {
        return "成功";
      }
      return "失败";
    },
     parseMarryKindType: function(value) {
      if (value == 1) {
        return "当前版本";
      }
      if (value == 2) {
        return "廉价版本";
      }
      return "";
    },
    serverTypeFilter: function(value) {
      return serverTypeList[value].name;
    },
    serverTagFilter: function(value) {
      return serverTagList[value].name;
    },
    serverStatusFilter: function(value) {
      return serverStatusList[value].name;
    }
  },
  data() {
    return {
      listLoading: false,
      tableKey: 0,
      total: 0,
      listQuery: {
        pageIndex: 1,
        centerServerName: "",
        platformId: undefined,
        serverType: undefined
      },
      textMap: {
        update: "编辑",
        create: "添加",
        refresh: "刷新"
      },
      serverTypeArray: [],
      serverTagArray: [],
      serverStatusArray: [],
      centerPlatformList: [],
      zhanChangServerList: [],

      dialogStatus: "",
      dialogPvVisible: false,
      dialogFormVisible: false,
      temp: {
        id: 0,
        centerPlatformId: undefined,
        centerServerId: undefined,
        successFlag: undefined,
        kindType: undefined,
        failMsg: undefined,
        updateTime: undefined
      },
      list: []
    };
  },
  created() {
    this.initData();
    // this.getList();
  },
  methods: {
    handleFilter: function() {
      this.listQuery.pageIndex = 1;
      console.log(this.listQuery)
      this.getList();
    },
    getList() {
      this.listLoading = true;
      getCenterPlatformMarrySetLogList(this.listQuery)
        .then(res => {
          this.list = res.itemArray;
          this.total = res.total;
          this.listLoading = false;
        })
        .catch(() => {
          this.listLoading = false;
        });
    },
    sendData(){
        updateCenterPlatformMarrySetSend(this.temp).then(res => {
          this.dialogPvVisible = false
          this.showSuccess()
          this.getList()
        })
        .catch(() => {
          
        });
    },
    handleCurrentChange(e) {
      this.listQuery.pageIndex = e;
      this.getList();
    },
    getPlatformName(value) {
      for (let i = 0, len = this.centerPlatformList.length; i < len; i++) {
        const item = this.centerPlatformList[i];
        if (item.centerPlatformId == value) {
          return item.centerPlatformName;
        }
      }
    },
    handleSend(e){
        this.dialogPvVisible = true;
        this.temp = Object.assign({}, e);
    },
    initData() {
      this.serverTypeArray = serverTypeList;
      this.serverTagArray = serverTagList;
      this.serverStatusArray = serverStatusList;
      getCenterPlatList().then(res => {
        this.centerPlatformList = res.itemArray;
      });
    },
    showSuccess() {
      this.$message({
        message: "修改成功",
        type: "success",
        duration: 1000
      });
    }
  }
};
</script>

