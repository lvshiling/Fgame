<template>
    <div>
        <div class="filter-container">
             <el-select v-model="listQuery.channelId" placeholder="渠道" style="width: 120px" class="filter-item" @change="handleChannelChange">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>

            <el-select v-model="listQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" clearable >
                <el-option v-for="item in tempPlatformList" :key="item.centerPlatformId" :label="item.platformName" :value="item.centerPlatformId" />
            </el-select>

            <div class="filter-item">
                <el-date-picker v-model="listQuery.timeArray" type="datetimerange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期">
                </el-date-picker>
            </div>
             <el-input placeholder="角色Id" v-model="listQuery.playerId" style="width: 200px;" class="filter-item"/>
             <el-input placeholder="兑换码" v-model="listQuery.code" style="width: 200px;" class="filter-item"/>
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
            <el-table-column label="ID" align="center" width="120px">
                <template slot-scope="scope">
                    <span>{{ scope.row.id }}</span>
                </template>
            </el-table-column>
            <el-table-column label="服务ID" align="center" width="100px">
                <template slot-scope="scope">
                    <span>{{ scope.row.serverId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="玩家角色ID" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.playerId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="兑换id" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.exchangeId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="兑换码" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.code }}</span>
                </template>
            </el-table-column>
            <el-table-column label="状态" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.status | parseFeedBackCenterStatus }}</span>
                </template>
            </el-table-column>
            <el-table-column label="金额" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.money }}</span>
                </template>
            </el-table-column>
            <el-table-column label="微信领取id" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.wxId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="订单id" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.orderId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="过期时间" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.expiredTime | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column label="创建时间" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.createTime | parseTime }}</span>
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
import { getCenterFeedBackFeeList } from "@/api/feedbackfee";
import { sdkTypeList, orderStateList } from "@/types/center";
import { parseTime } from "@/utils/index";
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getAllSdkType } from "@/api/center";
import { feedBackCenterStatus } from "@/types/feedback";
export default {
  name: "CenterOrderList",
  directives: {
    waves
  },
  filters: {
    parseSdkType: function(value) {
      if (sdkTypeList[value - 1]) {
        return sdkTypeList[value - 1].name;
      }
      return "";
    },
    parseTime: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
    },
    parseOrderStatus: function(value) {
      if (orderStateList[value]) {
        return orderStateList[value].name;
      }
      return "";
    },
    parseFeedBackCenterStatus: function(value) {
      if (feedBackCenterStatus[value]) {
        return feedBackCenterStatus[value].name;
      }
      return "";
    }
  },
  created() {
    this.initMetaData();
    this.skdTypeArray = sdkTypeList;
    getAllSdkType().then(res => {
      this.skdTypeArray = res.itemArray;
    });
  },
  data() {
    return {
      listLoading: false,
      tableKey: 0,
      total: 0,
      listQuery: {
        pageIndex: 1,
        sdkOrderId: undefined,
        orderId: undefined,
        mySdkType: undefined,
        timeArray: []
      },
      textMap: {
        update: "编辑",
        create: "添加"
      },
      dialogStatus: "",
      dialogPvVisible: false,
      dialogFormVisible: false,
      dialogRefreshVisible: false,
      temp: {
        sdkType: undefined,
        centerPlatformName: undefined
      },
      channelList: [],
      platformList: [],
      groupList: [],
      tempPlatformList: [],
      list: [],
      skdTypeArray: []
    };
  },
  methods: {
    handleFilter: function() {
      this.listQuery.pageIndex = 1;
      this.listQuery.sdkType = parseInt(this.listQuery.mySdkType);

      if (this.listQuery.timeArray && this.listQuery.timeArray.length == 2) {
        this.listQuery.startTime = this.listQuery.timeArray[0].valueOf();
        this.listQuery.endTime = this.listQuery.timeArray[1].valueOf();
      }
      console.log(this.listQuery);
      this.getList();
    },
    getList() {
      if (!this.listQuery.platformId) {
        this.$message({
          message: "请选择平台",
          type: "error",
          duration: 1000
        });
        return;
      }
      getCenterFeedBackFeeList(this.listQuery)
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
    },
    handleChannelChange: function(e) {
      if (e) {
        this.listQuery.sdkType = undefined;
        this.tempPlatformList = this.findPlatFormList(e);
        if (this.tempPlatFormList && this.tempPlatFormList.length > 0) {
          this.listQuery.sdkType = this.tempPlatFormList[0].sdkType;
        }
        this.listQuery.sdkType = undefined;
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
    },
    getSdkTypeName(value) {
      for (let i = 0, len = this.skdTypeArray.length; i < len; i++) {
        if (this.skdTypeArray[i].key == value) {
          return this.skdTypeArray[i].name;
        }
      }
      return value;
    }
  }
};
</script>

