<template>
    <div class="app-container">
        <div class="filter-container">
            <el-select v-model="listQuery.channelId" placeholder="渠道" style="width: 120px" class="filter-item" @change="handleChannelChange">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>

            <el-select v-model="listQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handlePlatformChange">
                <el-option v-for="item in tempPlatformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>

            <el-select v-model="listQuery.serverId" collapse-tags placeholder="服务器" clearable style="width: 220px" class="filter-item" >
              <el-option v-for="item in serverList" :key="item.id" :label="item.serverName" :value="item.id"/>
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
            <el-table-column  label="总充值金额" align="center" width="100px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.totalAmount }}</span>
                </template>
            </el-table-column>
            <el-table-column  label="总充值人数" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.totalPerson }}</span>
                </template>
            </el-table-column>
            <el-table-column  label="总注册人数" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.totalRegisterPerson }}</span>
                </template>
            </el-table-column>
            <el-table-column  label="今日充值金额" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.todayAmount}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="今日充值人数" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.todayPerson}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="今日注册人数" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.todayRegisterPerson}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="今日活跃人数" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.todayActivityPerson}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="三日活跃人数" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.threeDayActivityPerson}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="周活跃人数" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.weekActivityPerson}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="昨日充值金额" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.yestodayAmount}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="昨日充值人数" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.yestodayPerson}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="昨日注册人数" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.yestodayRegisterPerson}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="昨日活跃人数" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.yestodayAvtivityPerson}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="月充值金额" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.monthAmount}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="月充值人数" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.monthPerson}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="月活跃人数" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.monthActivityPerson}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="月ARPU" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.monthActivityPerson ==0 ? 0 : Math.round(scope.row.monthAmount/scope.row.monthActivityPerson*100)/100}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="月ARPPU" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.monthPerson ==0 ? 0 : Math.round(scope.row.monthAmount/scope.row.monthPerson*100)/100}}</span>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import { parseTime } from "@/utils/index";
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { Message, MessageBox } from "element-ui";
import { getGameOrderStatic } from "@/api/centerorder";
import { gameOrderMap } from "@/types/order";

export default {
  name: "GameOrderStatic",
  directives: {
    waves
  },
  filters: {
    parseTime: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
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
    parseOrderStatus: function(value) {
      if (gameOrderMap[value]) {
        return gameOrderMap[value].name;
      }
      return "";
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
        platformId: undefined,
        channelId: undefined,
        serverId: undefined,
        sdkType:undefined
      },
      textMap: {
        update: "编辑",
        create: "添加"
      },
      dialogStatus: "",
      dialogPvVisible: false,
      dialogFormVisible: false,
      temp: {},
      list: [],
      channelList: [],
      platformList: [],
      groupList: [],
      tempPlatformList: [],
      serverList: [],
      chatForbidTimeArray: [],
      monitorTemp: {}
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
      this.getList();
    },

    getList() {
      this.listLoading = true;
      getGameOrderStatic(this.listQuery)
        .then(res => {
          this.list = res.itemArray;
          this.total = res.total;
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
    handleSort(e) {
      if (!this.listQuery.serverId) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      this.listQuery.ordercol = parseInt(e.prop);
      this.listQuery.ordertype = 0;
      if (e.order == "descending") {
        this.listQuery.ordertype = 1;
      }
      this.getList();
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
    handleChannelChange: function(e) {
      if (e) {
        this.listQuery.platformId = undefined;
        this.tempPlatformList = this.findPlatFormList(e);
        if (this.tempPlatFormList && this.tempPlatFormList.length > 0) {
          this.listQuery.platformId = this.tempPlatFormList[0].platformId;
        }
        this.groupList = [];
        this.listQuery.serverId = undefined;
      }
    },
    handlePlatformChange: function(e) {
      console.log(e);
      if (e) {
        let item = this.findPlatFormItem(e);

        if (item) {
          this.listQuery.sdkType = item.sdkType;
          getCenterServerList(item.centerPlatformId).then(res => {
            this.serverList = res.itemArray;
            this.listQuery.serverId = undefined;
          });
        }else{
          this.listQuery.sdkType = undefined;
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

