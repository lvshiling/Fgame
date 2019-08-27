<template>
    <div class="app-container">
        <div class="filter-container">
            <el-select v-model="listQuery.channelId" placeholder="渠道" style="width: 120px" class="filter-item" @change="handleChannelChange">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>

            <el-select v-model="listQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handlePlatformChange">
                <el-option v-for="item in tempPlatformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>

            <el-select v-model="listQuery.serverId" collapse-tags placeholder="服务器" clearable style="width: 220px" class="filter-item" @change="handleServerChange" >
              <el-option v-for="item in serverList" :key="item.id" :label="item.serverName" :value="item.id"/>
            </el-select>

            <el-input placeholder="本地商品Id" v-model="listQuery.tradeId" style="width: 160px;" class="filter-item"/>
            <el-input placeholder="玩家Id" v-model="listQuery.playerId" style="width: 160px;" class="filter-item"/>
            <el-input placeholder="等级" v-model="listQuery.level" type="number" style="width: 160px;" class="filter-item"/>
            <el-select v-model="listQuery.state" placeholder="状态" style="width: 120px" class="filter-item" clearable >
                <el-option v-for="item in tradeItemStateArray" :key="item.key" :label="item.name" :value="item.key"/>
            </el-select>
            <div class="filter-item">
                <el-date-picker v-model="listQuery.timeArray" type="datetimerange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期">
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
            <el-table-column label="中心平台Id" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.platform }}</span>
                </template>
            </el-table-column>
            <el-table-column label="中心服务Id" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.serverId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="本地商品Id" align="center" width="160px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.tradeId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="玩家Id" align="center" width="160px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.playerId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="玩家名字" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.playerName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="物品Id" align="center" width="80px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.itemId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="物品数量" align="center" width="80px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.itemNum }}</span>
                </template>
            </el-table-column>
            <el-table-column label="等级" align="center" width="80px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.level }}</span>
                </template>
            </el-table-column>
            <el-table-column label="价格" align="center" width="80px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.gold }}</span>
                </template>
            </el-table-column>
            <el-table-column label="状态" align="center" width="100px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.status | parseStatus }}</span>
                </template>
            </el-table-column>
            <el-table-column label="购买者平台" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.buyPlayerPlatform }}</span>
                </template>
            </el-table-column>
            <el-table-column label="购买者服务" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.buyPlayerServerId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="购买者Id" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.buyPlayerServerId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="购买者名字" align="center" width="120px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.buyPlayerName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="创建时间" align="center" width="180px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.createTime | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column label="更新时间" align="center" width="180px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.updateTime | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column label="属性数据" align="center" min-width="380px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.propertyData }}</span>
                </template>
            </el-table-column>
        </el-table>
        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>
    </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import { parseTime } from "@/utils/index";
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { Message, MessageBox } from "element-ui";
import { getCenterTradeItemList } from "@/api/centertrade";
import { tradeItemStateList } from "@/types/tradeitem";

export default {
  name: "TradeItemList",
  directives: {
    waves
  },
  filters: {
    parseTime: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}");
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
    parseStatus:function(value){
        console.log(value)
        let item = tradeItemStateList[value]
        console.log(item)
        if(item){
            return item.name
        }
        return value
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
        timeArray: [],
        platformId: undefined,
        channelId: undefined,
        serverId: undefined,
        centerPlatformId: undefined,
        centerServerId: undefined,
        tradeId: undefined,
        playerId: undefined,
        level: undefined,
        state: undefined
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
      queryServerList: [],
      channelList: [],
      platformList: [],
      groupList: [],
      tempPlatformList: [],
      serverList: [],
      monitorTemp: {},
      tradeItemStateArray: []
    };
  },
  methods: {
    handleFilter: function() {
      if (!this.listQuery.platformId) {
        Message({
          message: "请选择平台",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      if (this.listQuery.timeArray && this.listQuery.timeArray.length == 2) {
        this.listQuery.startTime = this.listQuery.timeArray[0].valueOf();
        this.listQuery.endTime = this.listQuery.timeArray[1].valueOf();
      }
      this.list = [];
      this.getList();
    },

    getList() {
      this.listLoading = true;
      this.listQuery.centerPlatformId = undefined
      this.listQuery.centerServerId = undefined
      if(this.listQuery.platformId){
          let platformInfo = this.findPlatFormItem(this.listQuery.platformId)
          if(platformInfo){
              this.listQuery.centerPlatformId = platformInfo.centerPlatformId
          }
      }
      if(this.listQuery.serverId){
          let serverInfo = this.findServerItem(this.listQuery.serverId)
          if(serverInfo){
              this.listQuery.centerServerId = serverInfo.serverId
          }
      }
      getCenterTradeItemList(this.listQuery)
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
    initMetaData() {
      getAllChannel().then(res => {
        this.channelList = res.itemArray;
      });
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
        // this.tempPlatformList = this.platformList;
      });
      this.tradeItemStateArray = tradeItemStateList
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
          getCenterServerList(item.centerPlatformId).then(res => {
            this.serverList = res.itemArray;
            this.listQuery.serverId = undefined;
          });
        }
      }
    },
    handleServerChange: function(e) {
      if (e) {
        this.listQuery.serverId = e;
      } else {
        this.listQuery.serverId = undefined;
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

