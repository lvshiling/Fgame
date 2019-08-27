<template>
    <div>
          <div class="filter-container">
            <el-select v-model="jinyanListQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handleJinYanPlatformChange">
              <el-option v-for="item in platformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>
            <el-select v-model="jinyanListQuery.serverId" collapse-tags placeholder="服务器" clearable style="width: 220px" class="filter-item" @change="handleJinYanServerChange">
              <el-option v-for="item in jinyanServerList" :key="item.id" :label="item.serverName" :value="item.id"/>
            </el-select>
            <el-input v-model.trim="jinyanListQuery.playerName" placeholder="角色名" style="width: 200px;" class="filter-item" />
            <el-input v-model.trim="jinyanListQuery.reason" placeholder="禁言理由" style="width: 200px;" class="filter-item" />
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleJinYanFilter">搜索</el-button>
          </div>
        
        <el-table
          v-loading="listLoading"
          :key="jinyanTableKey"
          :data="jinyanUserList"
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
          <el-table-column label="禁言状态" width="250px" align="center">
            <template slot-scope="scope">
              <span v-if="scope.row.forbidChat == 1" style="color:#F56C6C">{{ scope.row.forbidChat | parseJin }}</span>
              <span v-else style="color:#67C23A">{{ scope.row.forbidChat | parseJin }}</span>
            </template>
          </el-table-column>
          <el-table-column label="禁言理由" width="200px">
            <template slot-scope="scope">
                <span style="color:#F56C6C;font-weight:bold;}">{{ scope.row.forbidChatText }}</span>
            </template>
          </el-table-column>
          <el-table-column label="禁言时间" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.forbidChatTime | parseTime }}</span>
            </template>
          </el-table-column>
          <el-table-column label="解禁时间" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.forbidChatEndTime | parseTimeSp }}</span>
            </template>
          </el-table-column>
          <el-table-column label="禁言者" min-width="120px">
            <template slot-scope="scope">
              <span>{{ scope.row.forbidChatName }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" width="260" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button v-if="scope.row.forbidChat == 1" type="primary" size="mini" @click="handleJieJinYan(scope.row)">解封</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="pagination-container" style="margin-top:15px;">
          <el-pagination :current-page="jinyanListQuery.pageIndex" :page-sizes="[20]" :total="jinyanUserCount" background layout="total, sizes, prev, pager, next, jumper" @current-change="handleJinYanCurrentChange"/>
        </div>

        <el-dialog :visible.sync="dialogUnForbidChatFormVisible" title="是否解禁用户">
          <el-form ref="dataForm" :model="unForbidChatTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="封禁用户名">
              <el-input v-model="unForbidChatTemp.playerName" :disabled="true"/>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogUnForbidChatFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateUnForbidChatPlayer">解禁</el-button>
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
  name: "JinYanPanel",
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

      // 禁言tab开始
      jinyanTableKey: 1,
      jinyanListQuery: {
        platformId: undefined,
        serverId: undefined,
        playerName: undefined,
        reason: undefined,
        pageIndex: 1,
        centerPlatformId: undefined,
        centerServerId: undefined
      },
      jinyanServerList: [],
      jinyanUserList: [],
      jinyanUserCount: 0,
      unForbidChatTemp: {},
      dialogUnForbidChatFormVisible: false
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
    /** ********************禁言设置********************* */
    // 禁言设置
    handleJinYanPlatformChange(e) {
      const item = this.findPlatFormItem(e);
      if (item) {
        getCenterServerList(item.centerPlatformId).then(res => {
          this.jinyanServerList = res.itemArray;
        });
      }
    },
    handleJinYanServerChange(e) {
      const serverInfo = this.findJinYanServerItem(e);
      if (serverInfo) {
        this.jinyanListQuery.centerPlatformId = serverInfo.centerPlatformId;
        this.jinyanListQuery.centerServerId = serverInfo.serverId;
      }
    },
    findJinYanServerItem(serverId) {
      const server = this.jinyanServerList.find(n => {
        return n.id == serverId;
      });
      if (server) {
        return server;
      }
      return undefined;
    },

    handleJinYanFilter(e) {
      // 搜索
      if (
        !this.jinyanListQuery.centerPlatformId ||
        !this.jinyanListQuery.centerServerId
      ) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      this.jinyanListQuery.pageIndex = 1;

      this.loadJinYan();
    },
    handleJinYanCurrentChange(e) {
      // 分页
      this.jinyanListQuery.pageIndex = e;
      this.loadJinYan();
    },
    handleJieJinYan(e) {
      // 解除禁言
      this.dialogUnForbidChatFormVisible = true;
      this.unForbidChatTemp = e;
      console.log(e);
    },
    updateUnForbidChatPlayer(e) {
      const postdata = {
        centerPlatformId: this.unForbidChatTemp.centerPlatformId,
        centerServerId: this.unForbidChatTemp.centerServerId,
        playerId: this.unForbidChatTemp.id
      };
      unForbidChatPlayer(postdata).then(res => {
        this.dialogUnForbidChatFormVisible = false;
        this.showSuccess();
        setTimeout(() => {
          this.loadJinYan();
        }, 300);
        // this.loadJinYan();
      });
    },
    loadJinYan() {
      this.listLoading = true;
      getJinYanPlayerList(this.jinyanListQuery).then(res => {
        this.jinyanUserList = res.itemArray;
        this.jinyanUserCount = res.total;
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
