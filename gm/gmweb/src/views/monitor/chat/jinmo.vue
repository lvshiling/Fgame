<template>
    <div>
          <div class="filter-container">
            <el-select v-model="jinMoListQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handleJinMoPlatformChange">
              <el-option v-for="item in platformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>
            <el-select v-model="jinMoListQuery.serverId" collapse-tags placeholder="服务器" clearable style="width: 220px" class="filter-item" @change="handleJinMoServerChange">
              <el-option v-for="item in jinMoServerList" :key="item.id" :label="item.serverName" :value="item.id"/>
            </el-select>
            <el-input v-model.trim="jinMoListQuery.playerName" placeholder="角色名" style="width: 200px;" class="filter-item" />
            <el-input v-model.trim="jinMoListQuery.reason" placeholder="禁默理由" style="width: 200px;" class="filter-item" />
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleJinMoFilter">搜索</el-button>
          </div>
        
        <el-table
          v-loading="listLoading"
          :key="jinMoTableKey"
          :data="jinMoUserList"
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
          <el-table-column label="禁默状态" width="250px" align="center">
            <template slot-scope="scope">
              <span v-if="scope.row.ignoreChat == 1" style="color:#F56C6C">{{ scope.row.ignoreChat | parseJin }}</span>
              <span v-else style="color:#67C23A">{{ scope.row.ignoreChat | parseJin }}</span>
            </template>
          </el-table-column>
          <el-table-column label="禁默理由" width="200px">
            <template slot-scope="scope">
                <span style="color:#F56C6C;font-weight:bold;}">{{ scope.row.ignoreChatText }}</span>
            </template>
          </el-table-column>
          <el-table-column label="禁默时间" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.ignoreChatTime | parseTime }}</span>
            </template>
          </el-table-column>
          <el-table-column label="解禁时间" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.ignoreChatEndTime | parseTimeSp }}</span>
            </template>
          </el-table-column>
          <el-table-column label="禁言者" min-width="120px">
            <template slot-scope="scope">
              <span>{{ scope.row.ignoreChatName }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" width="260" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button v-if="scope.row.ignoreChat == 1" type="primary" size="mini" @click="handleJieJinMo(scope.row)">解封</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="pagination-container" style="margin-top:15px;">
          <el-pagination :current-page="jinMoListQuery.pageIndex" :page-sizes="[20]" :total="jinMoUserCount" background layout="total, sizes, prev, pager, next, jumper" @current-change="handleJinMoCurrentChange"/>
        </div>

        <el-dialog :visible.sync="dialogUnJinMoChatFormVisible" title="是否解禁用户">
          <el-form ref="dataForm" :model="unJinMoChatTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="封禁用户名">
              <el-input v-model="unJinMoChatTemp.playerName" :disabled="true"/>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogUnJinMoChatFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateUnJinMoChatPlayer">解禁</el-button>
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
  name: "JinMoPanel",
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

      // 禁莫tab开始
      jinMoTableKey: 1,
      jinMoListQuery: {
        platformId: undefined,
        serverId: undefined,
        playerName: undefined,
        reason: undefined,
        pageIndex: 1,
        centerPlatformId: undefined,
        centerServerId: undefined
      },
      jinMoServerList: [],
      jinMoUserList: [],
      jinMoUserCount: 0,
      unJinMoChatTemp: {},
      dialogUnJinMoChatFormVisible: false,
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
    /** ********************禁默设置********************* */
    // 禁默设置
    handleJinMoPlatformChange(e) {
      const item = this.findPlatFormItem(e);
      if (item) {
        getCenterServerList(item.centerPlatformId).then(res => {
          this.jinMoServerList = res.itemArray;
        });
      }
    },
    handleJinMoServerChange(e) {
      const serverInfo = this.findJinMoServerItem(e);
      if (serverInfo) {
        this.jinMoListQuery.centerPlatformId = serverInfo.centerPlatformId;
        this.jinMoListQuery.centerServerId = serverInfo.serverId;
      }
    },
    findJinMoServerItem(serverId) {
      const server = this.jinMoServerList.find(n => {
        return n.id == serverId;
      });
      if (server) {
        return server;
      }
      return undefined;
    },

    handleJinMoFilter(e) {
      // 搜索
      if (
        !this.jinMoListQuery.centerPlatformId ||
        !this.jinMoListQuery.centerServerId
      ) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      this.jinMoListQuery.pageIndex = 1;

      this.loadJinMo();
    },
    handleJinMoCurrentChange(e) {
      // 分页
      this.jinMoListQuery.pageIndex = e;
      this.loadJinMo();
    },
    handleJieJinMo(e) {
      // 解除禁言
      this.dialogUnJinMoChatFormVisible = true;
      this.unJinMoChatTemp = e;
      console.log(e);
    },
    updateUnJinMoChatPlayer(e) {
      const postdata = {
        centerPlatformId: this.unJinMoChatTemp.centerPlatformId,
        centerServerId: this.unJinMoChatTemp.centerServerId,
        playerId: this.unJinMoChatTemp.id
      };
      console.log(postdata);
      unIgnoreChatPlayer(postdata).then(res => {
        this.dialogUnJinMoChatFormVisible = false;
        this.showSuccess();
        setTimeout(() => {
          this.loadJinMo();
        }, 300);
      });
    },
    loadJinMo() {
      this.listLoading = true;
      getJinMoPlayerList(this.jinMoListQuery).then(res => {
        this.jinMoUserList = res.itemArray;
        this.jinMoUserCount = res.total;
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
