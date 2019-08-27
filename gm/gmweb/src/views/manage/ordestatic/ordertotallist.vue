<template>
  <div class="app-container">
    <div class="filter-container">
      <el-select
        v-model="listQuery.channelId"
        placeholder="渠道"
        style="width: 120px"
        class="filter-item"
        @change="handleChannelChange"
      >
        <el-option
          v-for="item in channelList"
          :key="item.channelId"
          :label="item.channelName"
          :value="item.channelId"
        />
      </el-select>
      <el-select
        v-model="listQuery.platformId"
        placeholder="平台"
        style="width: 160px"
        class="filter-item"
        clearable
      >
        <el-option
          v-for="item in tempPlatformList"
          :key="item.sdkType"
          :label="item.platformName"
          :value="item.sdkType"
        />
      </el-select>
      <div class="filter-item">
        <el-date-picker
          v-model="listQuery.startEnd"
          type="datetimerange"
          range-separator="至"
          start-placeholder="开始时间"
          end-placeholder="结束时间"
        ></el-date-picker>
      </div>
      <el-button
        v-waves
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="handleFilter"
      >搜索</el-button>
        <!-- <el-button
        v-waves
        class="filter-item"
        type="primary"
        @click="handleDown"
      >下载</el-button> -->
      
    </div>

    <el-table
      v-loading="listLoading"
      :key="tableKey"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;margin-top:15px;"
    >
      <el-table-column label="SDK类型" align="left" min-width="150px">
        <template slot-scope="scope">
          <span>{{ getSdkTypeName(scope.row.sdkType) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="总充值金额" align="left" min-width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.orderMoney }}</span>
        </template>
      </el-table-column>
      <el-table-column label="总充值人数" align="left" min-width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.orderPlayerNum }}</span>
        </template>
      </el-table-column>
      <el-table-column label="订单总数" align="left" min-width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.orderNum }}</span>
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
import { Message, MessageBox } from "element-ui";
import { getCenterOrderStatic } from "@/api/centerorder";
import { sdkTypeList } from "@/types/center";
import { getOrderDatePlatformStatic } from '@/api/centerorder';
import { getAllSdkType } from "@/api/center";

export default {
  name: "OrderTotalList",
  directives: {
    waves
  },
  filters: {
    parseTime: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
    },
    parseSdkType: function(value) {
      if (sdkTypeList[value - 1]) {
        return sdkTypeList[value - 1].name;
      }
      if (value == -1) {
        return "合计";
      }
      return "";
    },
    parseSecond: function(value) {
      let hour = parseInt(value / 60 / 60 / 1000);
      let reseMinute = value % (60 * 60);
      let minute = parseInt(reseMinute / 60);
      let reseSecond = reseMinute % 60;
      return hour + "时" + minute + "分" + reseSecond + "秒";
    }
  },
  created() {
    this.initMetaData();
    // getOrderDatePlatformStatic().then(res => {
    // });
    // this.getList();
    getAllSdkType().then(res =>{
        this.skdTypeArray = res.itemArray
    });
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
        sdkType: undefined,
        startEnd:[],
        startTime:undefined,
        endTime:undefined,
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
      skdTypeArray: []
    };
  },
  methods: {
    handleFilter: function() {
       if (this.listQuery.startEnd && this.listQuery.startEnd.length == 2) {
        this.listQuery.startTime = this.listQuery.startEnd[0].valueOf();
        this.listQuery.endTime = this.listQuery.startEnd[1].valueOf();
      }
      if (this.listQuery.platformId==undefined || this.listQuery.platformId.length==0){
        this.listQuery.platformId = undefined
      }
      this.getList();
    },
   handleDown: function(){

   },
    getList() {
      this.listLoading = true;
      getOrderDatePlatformStatic(this.listQuery)
        .then(res => {
          this.list = res.itemArray;
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
    // handleSort(e) {
    //   if (!this.listQuery.serverId) {
    //     Message({
    //       message: "请选择服务器",
    //       type: "error",
    //       duration: 1.5 * 1000
    //     });
    //     return;
    //   }
    //   this.listQuery.ordercol = parseInt(e.prop);
    //   this.listQuery.ordertype = 0;
    //   if (e.order == "descending") {
    //     this.listQuery.ordertype = 1;
    //   }
    //   this.getList();
    // },
    initMetaData() {
      getAllChannel().then(res => {
        this.channelList = res.itemArray;
      });
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
        this.tempPlatformList = this.platformList;
      });
    },
    handleChannelChange: function(e) {
      if (e) {
        this.listQuery.sdkType = undefined;
        this.tempPlatformList = this.findPlatFormList(e);
        if (this.tempPlatFormList && this.tempPlatFormList.length > 0) {
          this.listQuery.sdkType = this.tempPlatFormList[0].sdkType;
        }
        this.groupList = [];
        this.listQuery.sdkType = undefined;
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
      if (value == -1) {
        return "合计";
      }
      return "";
    }
  }
};
</script>

