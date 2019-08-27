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
            <div class="filter-item">
                <el-date-picker v-model="listQuery.timeArray" type="datetimerange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期">
                </el-date-picker>
            </div>
             <el-input placeholder="最小金额" v-model="listQuery.minAmount" style="width: 150px;" class="filter-item" type="number"/>
             <el-input placeholder="最大金额" v-model="listQuery.maxAmount" style="width: 150px;" class="filter-item" type="number"/>
             <el-input placeholder="角色Id" v-model="listQuery.playerId" style="width: 200px;" class="filter-item"/>
             <el-input placeholder="账户Id" v-model="listQuery.userId" style="width: 200px;" class="filter-item"/>
             <el-input placeholder="订单号" v-model="listQuery.orderId" style="width: 200px;" class="filter-item"/>
             <el-input placeholder="平台订单号" v-model="listQuery.sdkOrderId" style="width: 200px;" class="filter-item"/>
             <el-input placeholder="角色名" v-model="listQuery.playerName" style="width: 200px;" class="filter-item"/>
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
        </div>

        <el-table
            v-loading="listLoading"
            :key="tableKey"
            :data="list"
            border
            fit
            highlight-current-row
            style="width: 100%;margin-top:15px;"
            @sort-change="handleSort">
            <el-table-column label="Id" align="center" width="100px" sortable="custom" prop="1">
                <template slot-scope="scope">
                    <span>{{ scope.row.id }}</span>
                </template>
            </el-table-column>
            <el-table-column label="创建时间" min-width="150px" align="left" sortable="custom" prop="10">
                <template slot-scope="scope">
                    <span>{{ scope.row.createTime | parseTime}}</span>
                </template>
            </el-table-column>
            <el-table-column label="金额" min-width="150px" align="left" sortable="custom" prop="8">
                <template slot-scope="scope">
                    <span>{{ scope.row.money}}</span>
                </template>
            </el-table-column>
            <el-table-column label="元宝" min-width="150px" align="left" sortable="custom" prop="13">
                <template slot-scope="scope">
                    <span>{{ scope.row.gold}}</span>
                </template>
            </el-table-column>
            <el-table-column label="角色id" min-width="150px" align="left" sortable="custom" prop="6">
                <template slot-scope="scope">
                    <span>{{ scope.row.playerId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="账号id" min-width="150px" align="left" sortable="custom" prop="5">
                <template slot-scope="scope">
                    <span>{{ scope.row.userId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="角色名" min-width="150px" align="left" sortable="custom" prop="14">
                <template slot-scope="scope">
                    <span>{{ scope.row.playerName}}</span>
                </template>
            </el-table-column>
            <el-table-column label="等级" min-width="150px" align="left" sortable="custom" prop="12">
                <template slot-scope="scope">
                    <span>{{ scope.row.playerLevel}}</span>
                </template>
            </el-table-column>
            <el-table-column label="订单号" min-width="150px" align="left" sortable="custom" prop="3">
                <template slot-scope="scope">
                    <span>{{ scope.row.orderId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="订单状态" min-width="150px" align="left" sortable="custom" prop="4">
                <template slot-scope="scope">
                    <span>{{ scope.row.orderStatus | parseOrderStatus}}</span>
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
import { getGameOrderList } from "@/api/centerorder";
import { gameOrderMap } from "@/types/order";

export default {
  name: "GameOrderList",
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
    parseOrderStatus : function(value){
        if(gameOrderMap[value]){
            return gameOrderMap[value].name
        }
        return ""
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
        allianceName: "",
        ordercol: 1,
        ordertype: 0,
        platformId: undefined,
        channelId: undefined,
        serverId: undefined
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
      this.listQuery.pageIndex = 1;
      this.listQuery.ordercol = 1;
      this.listQuery.ordertype = 0;
      this.listQuery.minAmount = parseInt(this.listQuery.minAmount);
      this.listQuery.maxAmount = parseInt(this.listQuery.maxAmount);
      if (this.listQuery.timeArray && this.listQuery.timeArray.length == 2) {
        this.listQuery.startTime = this.listQuery.timeArray[0].valueOf();
        this.listQuery.endTime = this.listQuery.timeArray[1].valueOf();
      }
      this.getList();
    },

    getList() {
      this.listLoading = true;
      getGameOrderList(this.listQuery)
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

