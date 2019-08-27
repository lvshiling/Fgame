<template>
<div class="app-container">
  <el-tabs type="border-card" v-model="activeName" @tab-click="handleClick">
    <el-tab-pane label="首充双倍重置" name="first">
        <el-form ref="dataForm" label-position="left" label-width="100px" style="width: 400px; margin-left:50px;">
            <el-form-item label="渠道">
                <el-select v-model="listQuery.channelId" placeholder="渠道" style="width: 120px" class="filter-item" @change="handleChannelChange">
                    <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
                </el-select>
            </el-form-item>
            <el-form-item label="平台">
                <el-select v-model="listQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handlePlatformChange">
                    <el-option v-for="item in tempPlatformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
                </el-select>
            </el-form-item>
            <el-form-item label="服务器">
                <el-select v-model="listQuery.serverId" collapse-tags placeholder="服务器" style="width: 220px" class="filter-item"  @change="handleServerChange">
                    <el-option v-for="item in serverList" :key="item.id" :label="item.serverName" :value="item.id"/>
                </el-select>
            </el-form-item>
            <el-form-item  v-if="listQuery.resetTime != undefined "  label="当前重置日期">
                <span>{{ listQuery.resetTime |  parseTime}}</span>
            </el-form-item>
            <el-button v-waves class="filter-item" type="primary" @click="handleSubmit">重置</el-button>    
        </el-form>
        <el-dialog :visible.sync="dialogFormVisible" title="是否确定重置">
          <div>是否确定重置</div>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateServer">确定</el-button>
          </div>
        </el-dialog>
    </el-tab-pane>
    <el-tab-pane label="首充双倍重置列表" name="second">
      <div class="filter-container">
            <el-select v-model="searchQuery.channelId" placeholder="渠道" style="width: 120px" class="filter-item" @change="handleSearchChannelChange">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>

            <el-select v-model="searchQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handleSearchPlatformChange">
                <el-option v-for="item in tempPlatformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>

            <el-select v-model="searchQuery.serverId" collapse-tags placeholder="服务器" clearable style="width: 180px" class="filter-item" >
              <el-option v-for="item in serverList" :key="item.id" :label="item.serverName" :value="item.id"/>
            </el-select>
               <el-table-column fixed="left" label="重置时间" align="center" min-width="100px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.createTime |  parseTime}}</span>
                </template>
            </el-table-column>
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
            <el-table-column fixed="left" label="日志id" align="center" min-width="100px"  >
                <template slot-scope="scope">
                    <span>{{ scope.row.id }}</span>
                </template>
            </el-table-column>
            <el-table-column fixed="left" label="服务器Id" align="center" min-width="100px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.serverId }}</span>
                </template>
            </el-table-column>
            <el-table-column fixed="left" label="重置时间" align="center" min-width="100px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.createTime |  parseTime}}</span>
                </template>
            </el-table-column>
        </el-table>

        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="searchQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>
    </el-tab-pane>
  </el-tabs>
</div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { Message, MessageBox } from "element-ui";
import { yesOrNoList } from "@/types/public";
import { parseTime } from "@/utils/index";
import {
  getServerDoubleCharge,
  resetServerDoubleCharge,
  serverDoubleChargeLoglist
} from "@/api/singleserver";


export default {
  name: "doublecharge",
  directives: {
    waves
  },
  filters: {
    parseTime: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}");
    }
  },
  data() {
    return {
      activeName: "first",
      listQuery: {
        channelId: undefined,
        platformId: undefined,
        serverId: undefined,
        resetTime: undefined
      },
      dialogFormVisible: false,
      channelList: [],
      platformList: [],
      tempPlatformList: [],
      serverList: [],
   
      //列表日志
      searchQuery:{
        channelId: undefined,
        platformId: undefined,
        serverId: undefined,
        pageIndex : 1
      },
      tableKey: 1,
      listLoading: false,
      list: [],
      total:0
    };
  },
  created() {
    this.initMetaData();
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    },
    handleSubmit(e) {
      if (!this.listQuery.serverId) {
        Message({
          message: "请选择参数",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      this.dialogFormVisible = true;
    },
    handleChannelChange: function(e) {
      if (e) {
        this.listQuery.platformId = undefined;
        this.tempPlatformList = this.findPlatFormList(e);
        if (this.tempPlatFormList && this.tempPlatFormList.length > 0) {
          this.listQuery.platformId = this.tempPlatFormList[0].platformId;
        }
        this.groupList = [];
        this.listQuery.serverId = undefined;
        this.listQuery.resetTime = undefined;
      }
    },
    handlePlatformChange: function(e) {
      console.log(e);
      if (e) {
        let item = this.findPlatFormItem(e);

        if (item) {
          getCenterServerList(item.centerPlatformId).then(res => {
            this.serverList = res.itemArray;
            this.listQuery.serverId = undefined;
             this.listQuery.resetTime = undefined;
          });
        }
      }
    },
    handleServerChange: function(e) {
      console.log(e);
      let myserverid = parseInt(e);
      var postData = {
        serverId: myserverid
      };
      getServerDoubleCharge(postData).then(res => {
          console.log(res);
        this.listQuery.resetTime = res.startTime;
      });
    },
    updateServer() {
      var postData = {
        serverId: this.listQuery.serverId,
      };
      resetServerDoubleCharge(postData).then(res => {
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    handleSearchChannelChange: function(e) {
      if (e) {
        this.searchQuery.platformId = undefined;
        this.tempPlatformList = this.findPlatFormList(e);
        if (this.tempPlatFormList && this.tempPlatFormList.length > 0) {
          this.searchQuery.platformId = this.tempPlatFormList[0].platformId;
        }
        this.groupList = [];
        this.searchQuery.serverId = undefined;
      }
    },
    handleSearchPlatformChange: function(e) {
      console.log(e);
      if (e) {
        let item = this.findPlatFormItem(e);

        if (item) {
          getCenterServerList(item.centerPlatformId).then(res => {
            this.serverList = res.itemArray;
            this.searchQuery.serverId = undefined;
          });
        }
      }
    },
    handleFilter: function() {
      if (!this.searchQuery.serverId) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      this.getList();
    },
    handleCurrentChange(e) {
      this.searchQuery.pageIndex = e;
      this.getList();
    },
    getList() {
      this.listLoading = true;
      serverDoubleChargeLoglist(this.searchQuery)
        .then(res => {
          this.list = res.itemArray;
          this.total = res.total;
          this.listLoading = false;
        })
        .catch(() => {
          this.listLoading = false;
        });
    },
    findPlatFormList(channelId) {
      if (!this.platformList || this.platformList.length == 0) {
        return;
      }
      return this.platformList.filter(function(item, index) {
        return item.channelId == channelId;
      });
    },
    findPlatFormItem(platformId) {
      let platform = this.platformList.find(n => {
        return n.platformId == platformId;
      });
      if (platform) {
        return platform;
      }
      return undefined;
    },
    findServerItem(id) {
      let server = this.serverList.find(n => {
        return n.id == id;
      });
      if (server) {
        return server;
      }
      return undefined;
    },
    initMetaData() {

      getAllChannel().then(res => {
        this.channelList = res.itemArray;
      });
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
        // this.tempPlatformList = this.platformList;
      });
    },
    showSuccess() {
      this.$message({
        message: "设置成功",
        type: "success",
        duration: 1000
      });
    }
  }
};
</script>


