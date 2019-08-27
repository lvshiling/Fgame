<template>
    <div class="app-container">
        <div class="filter-container">
            <el-select v-model="listQuery.channelId" placeholder="渠道" style="width: 120px" class="filter-item" @change="handleChannelChange">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>

            <el-select v-model="listQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handlePlatformChange">
                <el-option v-for="item in tpPlatFormList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>

            <el-select v-model="listQuery.serverId" collapse-tags placeholder="服务器" clearable style="width: 220px" class="filter-item" >
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
            :data="list"
            border
            fit
            highlight-current-row
            style="width: 100%;margin-top:15px;">
            <el-table-column label="日期" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.onLineDate | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column label="新增用户" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.num0 }}</span>
                </template>
            </el-table-column>
            <el-table-column :label="'留存率(%)'" min-width="120px">
                <el-table-column  label="1天后" align="center" width="100px" >
                    <template slot-scope="scope">
                        <span>{{ scope.row.num0 == 0 ? 0.00 : (scope.row.num1/scope.row.num0*100).toFixed(2) }}</span>
                    </template>
                </el-table-column>
                <el-table-column  label="2天后" align="center" width="100px" >
                    <template slot-scope="scope">
                        <span>{{ scope.row.num0 == 0 ? 0.00 : (scope.row.num2/scope.row.num0*100).toFixed(2) }}</span>
                    </template>
                </el-table-column>
                <el-table-column  label="3天后" align="center" width="100px" >
                    <template slot-scope="scope">
                        <span>{{ scope.row.num0 == 0 ? 0.00 : (scope.row.num3/scope.row.num0*100).toFixed(2) }}</span>
                    </template>
                </el-table-column>
                <el-table-column  label="4天后" align="center" width="100px" >
                    <template slot-scope="scope">
                        <span>{{ scope.row.num0 == 0 ? 0.00 : (scope.row.num4/scope.row.num0*100).toFixed(2) }}</span>
                    </template>
                </el-table-column>
                <el-table-column  label="5天后" align="center" width="100px" >
                    <template slot-scope="scope">
                        <span>{{ scope.row.num0 == 0 ? 0.00 : (scope.row.num5/scope.row.num0*100).toFixed(2) }}</span>
                    </template>
                </el-table-column>
                <el-table-column  label="6天后" align="center" width="100px" >
                    <template slot-scope="scope">
                        <span>{{ scope.row.num0 == 0 ? 0.00 : (scope.row.num6/scope.row.num0*100).toFixed(2) }}</span>
                    </template>
                </el-table-column>
                <el-table-column  label="7天后" align="center" width="100px" >
                    <template slot-scope="scope">
                        <span>{{ scope.row.num0 == 0 ? 0.00 : (scope.row.num7/scope.row.num0*100).toFixed(2) }}</span>
                    </template>
                </el-table-column>
                <el-table-column  label="14天后" align="center" width="100px" >
                    <template slot-scope="scope">
                        <span>{{ scope.row.num0 == 0 ? 0.00 : (scope.row.num14/scope.row.num0*100).toFixed(2) }}</span>
                    </template>
                </el-table-column>
                <el-table-column  label="30天后" align="center" width="100px" >
                    <template slot-scope="scope">
                        <span>{{ scope.row.num0 == 0 ? 0.00 : (scope.row.num30/scope.row.num0*100).toFixed(2) }}</span>
                    </template>
                </el-table-column>
            </el-table-column>
        </el-table>

        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>
    </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import permission from "@/directive/permission/index.js"; // 权限判断指令
import { getPlayerList } from "@/api/player";
import { parseTime } from "@/utils/index";
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { Message, MessageBox } from "element-ui";
import { getPlayerRetention } from "@/api/report";

export default {
  name: "PlayerRetentionList",
  directives: {
    waves,
    permission
  },
  filters: {
    parseTime: function(value) {
      return parseTime(value, "{y}-{m}-{d}");
    },
    parseSecond: function(value) {
      let hour = parseInt(value / 60 / 60 / 1000);
      let reseMinute = value % (60 * 60);
      let minute = parseInt(reseMinute / 60);
      let reseSecond = reseMinute % 60;
      return hour + "时" + minute + "分" + reseSecond + "秒";
    },
    parsesex: function(value) {
      if (value == 1) {
        return "男";
      }
      if (value == 2) {
        return "女";
      }
      return value;
    },
    parsePlayerRole: function(value) {
      let info = playerRoleMap[value - 1];
      if (info) {
        return info.name;
      }
      return "";
    },
    parseRetentionObject:function(value){
      let nowDate = new Date()
      nowDate.setDate(nowDate.getDate()-1)
      let nowTime = nowDate.getTime()
      let time = value.startTime
      if(time < nowTime){
        if(value.totalRegisterCount == 0){
          return 0.00
        }
        return (value.onLineCount/value.totalRegisterCount * 100).toFixed(2)
      }
      return ''
    }
  },
  created() {
    this.initMetaData();
    // this.getList();
  },
  data() {
    return {
      listLoading: false,
      tableKey: 0,
      total: 0,
      listQuery: {
        pageIndex: 1,
        beginTime: undefined,
        endTime: undefined,
        startEnd: [],
        platformId: undefined,
        channelId: undefined,
        serverId: undefined
      },
      dayArray:[],
      channelList: [],
      platformList: [],
      tpPlatFormList: [],
      groupList: [],
      serverList: [],
      list: []
    };
  },
  methods: {
    handleFilter: function() {
      if (!this.listQuery.serverId) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      console.log(this.listQuery);
      this.listQuery.pageIndex = 1;
      if (this.listQuery.startEnd && this.listQuery.startEnd.length == 2) {
        this.listQuery.beginTime = this.listQuery.startEnd[0].valueOf();
        this.listQuery.endTime = this.listQuery.startEnd[1].valueOf();
      }
      this.getList();
    },

    getList() {
      this.listLoading = true;
      getPlayerRetention(this.listQuery)
        .then(res => {
          this.list = res.itemArray;
          this.dayArray = res.dayArray;
          this.listLoading = false;
        })
        .catch(() => {
          this.listLoading = false;
        });
    },
    handleCurrentChange(e) {
      console.log(e);
      this.listQuery.pageIndex = e;
      this.getList();
    },
    initMetaData() {
      getAllChannel().then(res => {
        this.channelList = res.itemArray;
      });
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
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
    },
    handleChannelChange: function(e) {
      if (e) {
        this.listQuery.platformId = undefined;
        this.tpPlatFormList = this.findPlatFormList(e);
        this.groupList = [];
        this.listQuery.serverId = undefined;
      }
    },
    handlePlatformChange: function(e) {
      console.log(e);
      if (e) {
        let item = this.findPlatFormItem(e);

        if (item) {
          getCenterServerList(item.centerPlatformId).then(res => {
            this.listQuery.serverId = undefined;
            this.serverList = res.itemArray;
          });
        }
      }
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

