<template>
    <div class="app-container">
        <div class="filter-container">
            <el-select v-model="listQuery.channelId" placeholder="渠道" collapse-tags style="width: 160px" class="filter-item" clearable multiple @change="handleChannelChange">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>

            <el-select v-model="listQuery.platformId" placeholder="平台" collapse-tags style="width: 160px" class="filter-item" clearable multiple @change="handlePlatformChange">
                <el-option v-for="item in tempPlatformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>

            <el-select v-model="listQuery.serverId" collapse-tags placeholder="服务器" multiple clearable style="width: 220px" class="filter-item" @change="handleServerChange">
              <el-option v-for="item in serverList" :key="item.id" :label="item.serverName" :value="item.id"/>
            </el-select>
            <div class="filter-item" >
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
            :summary-method="getSummaries"
            show-summary
            style="width: 100%;margin-top:15px;">
            <el-table-column  label="日期" align="center" width="100px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.orderDate }}</span>
                </template>
            </el-table-column>
            <el-table-column  label="充值人数" width="150px" align="left" prop="orderPlayerNum">
                <template slot-scope="scope">
                    <span>{{ scope.row.orderPlayerNum }}</span>
                </template>
            </el-table-column>
            <el-table-column  label="充值次数" width="150px" align="left" prop="orderNum">
                <template slot-scope="scope">
                    <span>{{ scope.row.orderNum }}</span>
                </template>
            </el-table-column>
            <el-table-column  label="注册人数" width="150px" align="left" prop="totalPlayerCount">
                <template slot-scope="scope">
                    <span>{{ scope.row.totalPlayerCount }}</span>
                </template>
            </el-table-column>
            <el-table-column  label="新增注册人数" width="150px" align="left" prop="dateNewPlayerCount">
                <template slot-scope="scope">
                    <span>{{ scope.row.dateNewPlayerCount }}</span>
                </template>
            </el-table-column>
            <el-table-column  label="新增充值人数" width="150px" align="left" prop="orderFirstPlayerCount">
                <template slot-scope="scope">
                    <span>{{ scope.row.orderFirstPlayerCount }}</span>
                </template>
            </el-table-column>
            <el-table-column  label="新增首充金额" width="150px" align="left" prop="orderFirstMoneyCount">
                <template slot-scope="scope">
                    <span>{{ scope.row.orderFirstMoneyCount }}</span>
                </template>
            </el-table-column>
            <!-- <el-table-column  label="新增且充值人数" width="150px" align="left" prop="orderNewPlayerCount">
                <template slot-scope="scope">
                    <span>{{ scope.row.orderNewPlayerCount }}</span>
                </template>
            </el-table-column>
            <el-table-column  label="新增且充值金额" width="150px" align="left" prop="orderNewMoney">
                <template slot-scope="scope">
                    <span>{{ scope.row.orderNewMoney }}</span>
                </template>
            </el-table-column> -->
            <el-table-column  label="充值金额" width="150px" align="left" prop="orderMoney">
                <template slot-scope="scope">
                    <span>{{ scope.row.orderMoney}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="充值元宝" width="150px" align="left" prop="orderGold">
                <template slot-scope="scope">
                    <span>{{ scope.row.orderGold}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="ARPPU" width="150px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.orderPlayerNum ==0 ? 0 : Math.round(scope.row.orderMoney/scope.row.orderPlayerNum*100)/100}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="ARPU" width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.totalPlayerCount ==0 ? 0 : Math.round(scope.row.orderMoney/scope.row.totalPlayerCount*100)/100}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="服务器数量" width="150px" align="left" prop="serverCount">
                <template slot-scope="scope">
                    <span>{{ scope.row.serverCount}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="单服平均充值" width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.serverCount ==0 ? 0 : Math.round(scope.row.orderMoney/scope.row.serverCount*100)/100}}</span>
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
import { getAllUserCenterServerList } from "@/api/center";
import { Message, MessageBox } from "element-ui";
import { getCenterOrderDateStatic } from "@/api/centerorder";
import { gameOrderMap } from "@/types/order";

export default {
  name: "CenterOrderDateStatic",
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
        platformId: [],
        channelId: [],
        serverId: [],
        startEnd: [],
        sdkType: undefined
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
      allserverList: [],
      groupList: [],
      tempPlatformList: [],
      tempPlatformIdList: [],
      serverList: [],
      serverIdList: [],
      chatForbidTimeArray: [],
      monitorTemp: {},
      prePlatformAllSelected: false,
      preServerAllSelected: false
    };
  },
  methods: {
    handleFilter: function() {
      if(!this.listQuery.channelId){
        Message({
          message: "请选渠道",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      if (this.listQuery.startEnd && this.listQuery.startEnd.length == 2) {
        this.listQuery.startTime = this.listQuery.startEnd[0].valueOf();
        this.listQuery.endTime = this.listQuery.startEnd[1].valueOf();
      }
      console.log(this.listQuery);
      this.getList();
    },

    getList() {
      this.listLoading = true;
      getCenterOrderDateStatic(this.listQuery)
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

      getAllChannel().then(res => {
        this.channelList = res.itemArray;
      });
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
        // this.tempPlatformList = this.platformList;
      });
      getAllUserCenterServerList().then(res => {
        this.allserverList = res.itemArray;
      });
    },
    handleChannelChange: function(e) {
      this.listQuery.platformId = [];
      this.listQuery.serverId = [];
      this.prePlatformAllSelected = false;
      this.preServerAllSelected = false;
      this.tempPlatformList = this.findPlatFormList(e);
      if (this.tempPlatformList.length > 0) {
        let allSelect = {
          centerPlatformId: -1,
          channelId: -1,
          platformId: -1,
          platformName: "全选",
          sdkType: -1
        };
        this.tempPlatformList.unshift(allSelect);
      }
      this.tempPlatformIdList = this.getPlatformIdList(this.tempPlatformList);
      this.listQuery.platformId = [];
    },
    handlePlatformChange: function(e) {
      let allFlag = this.findAllSelect(e);
      if (allFlag) {
        if (!this.prePlatformAllSelected) {
          this.listQuery.platformId = this.tempPlatformIdList;
          this.prePlatformAllSelected = true;
        } else {
          if (e.length != this.tempPlatformIdList.length) {
            const index = this.listQuery.platformId.indexOf(-1);
            this.listQuery.platformId.splice(index, 1); // 排除全选选项
            this.prePlatformAllSelected = false;
          }
        }
      } else {
        if (this.prePlatformAllSelected) {
          console.log("选中后");
          this.listQuery.platformId = [];
          this.prePlatformAllSelected = false;
        }
      }
      this.serverList = this.findServerList(this.listQuery.platformId);
      if (this.serverList.length > 0) {
        let allSelect = {
          centerPlatformId: -1,
          id: -1,
          serverId: -1,
          serverName: "全选",
          serverType: -1
        };
        this.serverList.unshift(allSelect);
      }
      this.serverIdList = this.getServerIdList(this.serverList);
      this.preServerAllSelected = false;
      this.listQuery.serverId = [];
    },
    handleServerChange: function(e) {
      let allFlag = this.findAllSelect(e);
      if (allFlag) {
        if (!this.preServerAllSelected) {
          this.listQuery.serverId = this.serverIdList;
          this.preServerAllSelected = true;
        } else {
          if (e.length != this.serverIdList.length) {
            const index = this.listQuery.serverId.indexOf(-1);
            this.listQuery.serverId.splice(index, 1); // 排除全选选项
            this.preServerAllSelected = false;
          }
        }
      } else {
        if (this.preServerAllSelected) {
          this.listQuery.serverId = [];
          this.preServerAllSelected = false;
        }
      }
      console.log(this.listQuery.serverId);
    },
    findPlatFormList(channelId) {
      if (!this.platformList || this.platformList.length == 0) {
        return;
      }
      return this.platformList.filter(function(item, index) {
        if (Array.isArray(channelId)) {
          for (let i = 0, len = channelId.length; i < len; i++) {
            if (channelId[i] == item.channelId) {
              return true;
            }
          }
          return false;
        }
        return item.channelId == channelId;
      });
    },
    findServerList(platformList) {
      if (!this.allserverList || this.allserverList.length == 0) {
        return;
      }
      this.serverList = [];
      let centerPlatformList = [];
      if (Array.isArray(platformList)) {
        for (let i = 0, len = platformList.length; i < len; i++) {
          let item = this.findPlatFormItem(platformList[i]);
          if (item && item.centerPlatformId) {
            centerPlatformList.push(item.centerPlatformId);
          }
        }
      } else {
        let item = this.findPlatFormItem(platformList);
        if (item && item.centerPlatformId) {
          centerPlatformList.push(item.centerPlatformId);
        }
      }
      if (centerPlatformList.length > 0) {
        return this.allserverList.filter(function(item, index) {
          for (let i = 0, len = centerPlatformList.length; i < len; i++) {
            if (centerPlatformList[i] == item.centerPlatformId) {
              return true;
            }
          }
          return false;
        });
      }
      return [];
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
    findAllSelect(arrayList) {
      if (Array.isArray(arrayList)) {
        for (let i = 0, len = arrayList.length; i < len; i++) {
          if (arrayList[i] == -1) {
            return true;
          }
        }
      }
      return false;
    },
    getPlatformIdList(platformList) {
      let rst = [];
      for (let i = 0, len = platformList.length; i < len; i++) {
        rst.push(platformList[i].platformId);
      }
      return rst;
    },
    getServerIdList(serverList) {
      let rst = [];
      for (let i = 0, len = serverList.length; i < len; i++) {
        rst.push(serverList[i].id);
      }
      return rst;
    },
    showSuccess() {
      this.$message({
        message: "设置成功",
        type: "success",
        duration: 1000
      });
    },
    getSummaries(param) {
      const { columns, data } = param;
      const sums = [];
      columns.forEach((column, index) => {
        if (index === 0) {
          sums[index] = "合计";
          return;
        }
        const values = data.map(item => Number(item[column.property]));
        if (index === 11) {
          sums[index] = values[values.length - 1];
          return;
        }
        sums[index] = values.reduce((prev, curr) => {
          const value = Number(curr);
          if (!isNaN(value)) {
            return prev + curr;
          } else {
            return prev;
          }
        }, 0);
        sums[9] = sums[1]==0 ? 0 : Math.round(sums[7]/sums[1]*100)/100
        sums[10] = sums[3]==0 ? 0 : Math.round(sums[7]/sums[3]*100)/100
        sums[12] = sums[11]==0 ? 0 : Math.round(sums[7]/sums[11]*100)/100
        // if (values.every(value => isNaN(value))) {
        //   sums[index] = values.reduce((prev, curr) => {
        //     const value = Number(curr);
        //     if (!isNaN(value)) {
        //       return prev + curr;
        //     } else {
        //       return prev;
        //     }
        //   }, 0);
        //   // sums[index] += " 元";
        // } else {
        //   sums[index] = "N/A";
        // }
      });

      return sums;
    }
  }
};
</script>

