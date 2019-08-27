<template>
    <div>
        <div class="filter-container">
            <el-select v-model="listQuery.channelId" placeholder="渠道" clearable style="width: 120px" class="filter-item" @change="handleChannelChange">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>

            <el-select v-model="listQuery.platformId" placeholder="平台" clearable style="width: 160px" class="filter-item" @change="handlePlatformChange">
                <el-option v-for="item in tempPlatformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>

            <el-select v-model="listQuery.serverId" collapse-tags placeholder="服务器" clearable style="width: 220px" class="filter-item" >
              <el-option v-for="item in serverList" :key="item.id" :label="item.serverName" :value="item.id"/>
            </el-select>

            <el-input placeholder="玩家名" v-model="listQuery.playerName" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
            <el-input placeholder="玩家id" v-model="listQuery.playerId" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
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
            <el-table-column label="玩家Id" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.playerId }}</span>
                </template>
            </el-table-column>
           <el-table-column label="角色名" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.playerName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="渠道" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ findChannelName(scope.row.channelId) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="平台" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ findPlatformName(scope.row.platformId) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="服务器名" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.serverName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="扶持元宝" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.gold }}</span>
                </template>
            </el-table-column>
            <el-table-column label="扶持人" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.userName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="扶持时间" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.chargeTime | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column label="扶持原因" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.reason }}</span>
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
import {
  supportPlayerLog
} from "@/api/supportplayer";
import { parseTime } from "@/utils/index";
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { Message, MessageBox } from "element-ui";
import { chatForbidTimeList } from "@/types/chat";
import { playerSupportMap } from "@/types/player";
import { supportAmount } from "@/types/manage";

export default {
  name: "SupportPlayerLog",
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
    // this.getList();
  },
  data() {
    return {
      listLoading: false,
      tableKey: 0,
      total: 0,
      listQuery: {
        pageIndex: 1,
        playerName: "",
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
      this.getList();
    },

    getList() {
      this.listLoading = true;
      supportPlayerLog(this.listQuery)
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
        })
        if(list && list.length > 0){
            return list[0].channelName
        }
        return ""
    },
    findPlatformName(platformId) {
        let item = this.findPlatFormItem(platformId)
        if(item){
            return item.platformName
        }
        return ""
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

