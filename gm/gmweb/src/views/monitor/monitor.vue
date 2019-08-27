<template>
    <div class="app-container">
        <div class="filter-container">
            <el-input placeholder="渠道名" v-model="listQuery.channelName" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleScoket">连接socket</el-button>
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleSendAuth">队列触发</el-button>

            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleSensitive">敏感词测试</el-button>
        </div>
    </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import basemessage from "@/proto/basic_pb";
import messages from "@/proto/login_pb";
import commonpb from "@/proto/common_pb";
import messagetype from "@/proto/messagetype_pb";
import { getToken } from "@/utils/auth";
import { newWebSocket } from "@/websocket/websocket";
import { QueueDelegate } from "@/websocket/queuedelegate";
import { makeSensitiveMap, replaceSensitiveWord } from "@/utils/sensitive";
export default {
  name: "MonitorList",
  directives: {
    waves
  },
  created() {},

  data() {
    return {
      listLoading: false,
      chatKey: 0,
      total: 0,
      listQuery: {
        pageIndex: 1,
        channelName: ""
      },

      textMap: {
        update: "编辑",
        create: "添加"
      },
      dialogStatus: "",
      dialogPvVisible: false,
      dialogFormVisible: false,
      temp: {},
      list: []
    };
  },

  methods: {
    handleSensitive: function() {
      let minganci = "小毛驴,小,1,2,3,4,5,6,7,8,9,好";
      this.minGanCiList = minganci.split(",");
      // for(const item of this.minGanCiList){
      //   console.log(item)
      // }
      this.sensitiveMap = makeSensitiveMap(this.minGanCiList);
      console.log(this.minGanCiList);
      console.log(this.sensitiveMap);
      let chatMsg =
        "我有一只小毛驴从来也不骑，1，2，3，4，5，6，哎呀呀，22223321,好久没吃hi";
      chatMsg = replaceSensitiveWord(
        this.sensitiveMap,
        chatMsg,
        '<span style="background:#E6A23C">',
        "</span>"
      );
      console.log(chatMsg);
    },
    handleFilter: function() {
      this.listQuery.pageIndex = 1;
      let login = new messages.CGLogin();
      login.setPlayerid(3);
      login.setToken("666");
      console.log(login);
      var mess = new basemessage.Message();
      mess.setMessagetype(messagetype.QiPaiMessageType.GCLOGINTYPE);
      mess.setExtension(messages.cglogin, login);
      console.log(mess);

      let byteinfo = mess.serializeBinary();
      console.log(byteinfo);

      let nexinfo = basemessage.Message.deserializeBinary(byteinfo);
      console.log(nexinfo);

      let dataInfo = nexinfo.getExtension(messages.cglogin);
      console.log(dataInfo);
      console.log("解析出来后的token");
      console.log(dataInfo.getToken());

      let errorCode = commonpb.ErrorCode.ROOMNOEXIST;
      console.log(errorCode);
    },
    handleScoket: function() {
      this.queuedelegate = new QueueDelegate();
      this.websock = newWebSocket(
        "ws://localhost:9090/websocket",
        this.queuedelegate
      );
    },
    handleSendAuth() {
      this.queuedelegate = new QueueDelegate();
      this.queuedelegate.begin();
    }
  }
};
</script>

