<template>
    <div>
          <div class="filter-container">
            <el-select v-model="fengJinListQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handleFengjinPlatformChange">
              <el-option v-for="item in platformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>
            <el-select v-model="fengJinListQuery.serverId" collapse-tags placeholder="服务器" clearable style="width: 220px" class="filter-item" @change="handleFengJinServerChange">
              <el-option v-for="item in fengjinServerList" :key="item.id" :label="item.serverName" :value="item.id"/>
            </el-select>
            <el-input v-model.trim="fengJinListQuery.playerName" placeholder="角色名" style="width: 200px;" class="filter-item" />
            <el-input v-model.trim="fengJinListQuery.reason" placeholder="封禁理由" style="width: 200px;" class="filter-item" />
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFengJinFilter">搜索</el-button>
          </div>
        <el-table
          v-loading="listLoading"
          :key="fengjinTableKey"
          :data="fengjinUserList"
          border
          fit
          highlight-current-row
          style="width: 100%;margin-top:15px;">
          <el-table-column label="角色ID" align="center" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.id }}</span>
            </template>
          </el-table-column>
          <el-table-column label="账户ID" align="center" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.playerId }}</span>
            </template>
          </el-table-column>
          <el-table-column label="角色名" align="center" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.playerName }}</span>
            </template>
          </el-table-column>
          <el-table-column label="封禁状态" width="250px" align="center">
            <template slot-scope="scope">
              <span v-if="scope.row.forbid == 1" style="color:#F56C6C">{{ scope.row.forbid | parseJin }}</span>
              <span v-else style="color:#67C23A">{{ scope.row.forbid | parseJin }}</span>
            </template>
          </el-table-column>
          <el-table-column label="封禁理由" width="200px">
            <template slot-scope="scope">
              <span style="color:#F56C6C;font-weight:bold;}">{{ scope.row.forbidText }}</span>
            </template>
          </el-table-column>
          <el-table-column label="封禁时间" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.forbidTime | parseTime }}</span>
            </template>
          </el-table-column>
          <el-table-column label="解禁时间" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.forbidEndTime | parseTimeSp }}</span>
            </template>
          </el-table-column>
          <el-table-column label="封禁者" min-width="120px">
            <template slot-scope="scope">
              <span>{{ scope.row.forbidName }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" width="260" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button v-if="scope.row.forbid == 1" type="primary" size="mini" @click="handleJieFengJin(scope.row)">解封</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="pagination-container" style="margin-top:15px;">
          <el-pagination :current-page="fengJinListQuery.pageIndex" :page-sizes="[20]" :total="fengjinUserCount" background layout="total, sizes, prev, pager, next, jumper" @current-change="handleFengJinCurrentChange"/>
        </div>

        <el-dialog :visible.sync="dialogUnForbidFormVisible" title="是否解禁用户">
          <el-form ref="dataForm" :model="unForbidTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="封禁用户名">
              <el-input v-model="unForbidTemp.playerName" :disabled="true"/>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogUnForbidFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateUnForbidPlayer">解禁</el-button>
          </div>
        </el-dialog>
    </div>
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { getSensitive, saveUserInfo } from "@/api/sensitive";
import { parseTime } from "@/utils/index";
import {
  getFengJinPlayerList,
  getJinYanPlayerList,
  forbidPlayer,
  unForbidPlayer,
  forbidChatPlayer,
  unForbidChatPlayer,
  ignoreChatPlayer,
  unIgnoreChatPlayer,
  getJinMoPlayerList
} from "@/api/player";
import { chatTypeList, chatMethodList, chatForbidTimeList } from "@/types/chat";
import { makeSensitiveMap, replaceSensitiveWord } from "@/utils/sensitive";
import basemessage from "@/proto/basic_pb";
import chat_pb from "@/proto/chat_pb";
import chatmessagetype_pb from "@/proto/chatmessagetype_pb";
import { Message, MessageBox } from "element-ui";

export default {
  name: "FengJinPanel",
  directives: {
    waves
  },
  filters: {
    parseTime: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
    },
    parseTimeSp: function(value) {
      if (!value) {
        return "永久";
      }
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
    },
    parseJin: function(value) {
      if (value == 1) {
        return "已封禁";
      }
      if (value == 0) {
        return "已解封";
      }
    },
    chatTypeFilter: function(value) {
      return chatTypeList[value].name;
    },
    chatMethodFilter: function(value) {
      return chatMethodList[value].name;
    }
  },
  data() {
    return {
      listLoading: false,
      chatKey: 0,
      total: 0,
      dialogFormVisible: false,
      dialogForbidFormVisible: false,
      dialogForbidChatFormVisible: false,
      dialogIgnoreChatFormVisible: false,
      websocket: undefined,
      platformList: [],
      serverList: [],
      minGanCiList: [],
      sensitiveMap: undefined,
      chatTypeList: [],
      chatList: [],
      chatState: false,

      // 封禁tab开始
      fengjinTableKey: 1,
      fengJinListQuery: {
        platformId: undefined,
        serverId: undefined,
        playerName: undefined,
        reason: undefined,
        pageIndex: 1,
        centerPlatformId: undefined,
        centerServerId: undefined
      },
      fengjinServerList: [],
      fengjinUserList: [],
      fengjinUserCount: 0,
      unForbidTemp: {}, // 解禁用户传入对象
      dialogUnForbidFormVisible: false
    };
  },
  created() {
    this.initMetaData();
  },
  mounted() {
    //   this.websocket.close()
  },
  beforeDestroy() {},
  destroyed() {},
  methods: {
    initMetaData() {
      this.chatTypeList = chatTypeList;
      this.chatForbidTimeArray = chatForbidTimeList;
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
      });
    },
    findPlatFormItem(platformId) {
      const platform = this.platformList.find(n => {
        return n.platformId == platformId;
      });
      if (platform) {
        return platform;
      }
      return undefined;
    },
    /** **********************************封禁设置 ****************************/

    // 封禁设置
    handleFengjinPlatformChange(e) {
      const item = this.findPlatFormItem(e);
      if (item) {
        getCenterServerList(item.centerPlatformId).then(res => {
          this.fengjinServerList = res.itemArray;
        });
      }
    },
    handleFengJinServerChange(e) {
      const serverInfo = this.findFengJinServerItem(e);
      if (serverInfo) {
        this.fengJinListQuery.centerPlatformId = serverInfo.centerPlatformId;
        this.fengJinListQuery.centerServerId = serverInfo.serverId;
      }
    },
    findFengJinServerItem(serverId) {
      const server = this.fengjinServerList.find(n => {
        return n.id == serverId;
      });
      if (server) {
        return server;
      }
      return undefined;
    },

    handleFengJinFilter(e) {
      // 搜索
      if (
        !this.fengJinListQuery.centerPlatformId ||
        !this.fengJinListQuery.centerServerId
      ) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      this.fengJinListQuery.pageIndex = 1;
      console.log(this.fengJinListQuery);

      this.loadFengJin();
    },
    handleFengJinCurrentChange(e) {
      // 分页
      this.fengJinListQuery.pageIndex = e;
      this.loadFengJin();
    },
    handleJieFengJin(e) {
      this.dialogUnForbidFormVisible = true;
      this.unForbidTemp = e;
      console.log("解除封禁");
      console.log(e);
      // 解除封禁
    },
    updateUnForbidPlayer(e) {
      const postdata = {
        centerPlatformId: this.unForbidTemp.centerPlatformId,
        centerServerId: this.unForbidTemp.centerServerId,
        playerId: this.unForbidTemp.id
      };
      unForbidPlayer(postdata).then(res => {
        this.dialogUnForbidFormVisible = false;
        this.showSuccess();
        setTimeout(() => {
          this.loadFengJin();
        }, 500);
        // this.loadFengJin();
      });
    },
    loadFengJin() {
      this.listLoading = true;
      getFengJinPlayerList(this.fengJinListQuery).then(res => {
        this.fengjinUserList = res.itemArray;
        this.fengjinUserCount = res.total;
        this.listLoading = false;
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
