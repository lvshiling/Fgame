<template>
    <div>
        <div class="filter-container">
            <el-input placeholder="玩家id" disabled v-model="listQuery.playerId" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
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
            <el-table-column label="Id" align="center" width="100px">
                <template slot-scope="scope">
                    <span>{{ scope.row.id }}</span>
                </template>
            </el-table-column>
           <el-table-column label="玩家Id" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.playerId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="是否读取" align="center" width="80px">
                <template slot-scope="scope">
                    <span>{{ scope.row.isRead }}</span>
                </template>
            </el-table-column>
            <el-table-column label="是否已领取附件" align="center" width="80px">
                <template slot-scope="scope">
                    <span>{{ scope.row.isGetAttachment }}</span>
                </template>
            </el-table-column>
            <el-table-column label="标题" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.title }}</span>
                </template>
            </el-table-column>
            <el-table-column label="内容" align="center" min-width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.content }}</span>
                </template>
            </el-table-column>
            <el-table-column label="附件信息" align="center" min-width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.attachementInfo }}</span>
                </template>
            </el-table-column>
            <el-table-column label="创建时间" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.createTime | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column label="删除时间" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.deleteTime | parseTime }}</span>
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
import { getPlayerMail } from "@/api/playerinfo";
import { parseTime } from "@/utils/index";
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { Message, MessageBox } from "element-ui";
import { chatForbidTimeList } from "@/types/chat";
import { playerSupportMap } from "@/types/player";
import { supportAmount } from "@/types/manage";

export default {
  name: "PlayerEMail",
  directives: {
    waves
  },
  props: ["serverId", "playerId"],
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
    parseFuchiType: function(value) {
      let item = playerSupportMap[value];
      if (item) {
        return item.name;
      }
      return "";
    }
  },
  created() {
    this.initMetaData();
  },
  data() {
    return {
      paramPlayerId: this.playerId,
      paramServerId: this.serverId,
      listLoading: false,
      tableKey: 0,
      total: 0,
      listQuery: {
        pageIndex: 1,
        playerName: "",
        ordercol: 1,
        ordertype: 0,
        startEnd: [],
        serverId: 2,
        beginTime: undefined,
        endTime: undefined,
        platformId: undefined,
        channelId: undefined
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
      dialogUnForbidFormVisible: false,
      dialogFuChiFormVisible: false,
      channelList: [],
      platformList: [],
      groupList: [],
      tempPlatformList: [],
      serverList: [],
      chatForbidTimeArray: [],
      playerSupportList: [],
      fuchiInfo: {
        playerId: undefined,
        gold: undefined,
        privilege: undefined
      },
      supportAmountArray: []
    };
  },
  methods: {
    handleFilter: function() {
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
      if (this.listQuery.startEnd && this.listQuery.startEnd.length == 2) {
        this.listQuery.beginTime = this.listQuery.startEnd[0].valueOf();
        this.listQuery.endTime = this.listQuery.startEnd[1].valueOf();
      }
      getPlayerMail(this.listQuery)
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
      this.listQuery.playerId = this.paramPlayerId;
      this.listQuery.serverId = this.paramServerId;
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
    findChannelName(channelId) {
      let list = this.channelList.filter(function(item, index) {
        return item.channelId == channelId;
      });
      if (list && list.length > 0) {
        return list[0].channelName;
      }
      return "";
    },
    findPlatformName(platformId) {
      let item = this.findPlatFormItem(platformId);
      if (item) {
        return item.platformName;
      }
      return "";
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

