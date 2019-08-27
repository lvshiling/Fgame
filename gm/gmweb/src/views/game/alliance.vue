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
        @change="handlePlatformChange"
      >
        <el-option
          v-for="item in tempPlatformList"
          :key="item.platformId"
          :label="item.platformName"
          :value="item.platformId"
        />
      </el-select>

      <el-select
        v-model="listQuery.serverId"
        collapse-tags
        placeholder="服务器"
        clearable
        style="width: 220px"
        class="filter-item"
      >
        <el-option
          v-for="item in serverList"
          :key="item.id"
          :label="item.serverName"
          :value="item.id"
        />
      </el-select>

      <el-input
        placeholder="仙盟名"
        v-model="listQuery.allianceName"
        style="width: 200px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"
      />

      <el-button
        v-waves
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="handleFilter"
      >搜索</el-button>
    </div>

    <el-table
      v-loading="listLoading"
      :key="tableKey"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;margin-top:15px;"
      @sort-change="handleSort"
    >
      <el-table-column
        fixed="left"
        label="仙盟Id"
        align="center"
        width="180px"
        sortable="custom"
        prop="1"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column
        fixed="left"
        label="仙盟名称"
        width="150px"
        align="left"
        sortable="custom"
        prop="2"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.allianceName}}</span>
        </template>
      </el-table-column>
      <el-table-column
        fixed="left"
        label="等级"
        width="120px"
        align="left"
        sortable="custom"
        prop="3"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.allianceLevel}}</span>
        </template>
      </el-table-column>
      <el-table-column
        fixed="left"
        label="总战斗力"
        width="200px"
        align="left"
        sortable="custom"
        prop="4"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.totalForce}}</span>
        </template>
      </el-table-column>
      <el-table-column
        fixed="left"
        label="玩家数"
        width="120px"
        align="left"
        sortable="custom"
        prop="6"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.playerCount}}</span>
        </template>
      </el-table-column>
      <el-table-column
        fixed="left"
        label="仙盟公告"
        min-width="150px"
        align="left"
        sortable="custom"
        prop="7"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.notice}}</span>
        </template>
      </el-table-column>
      <el-table-column
        fixed="left"
        label="创建时间"
        width="180px"
        align="left"
        sortable="custom"
        prop="5"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.createTime | parseTime}}</span>
        </template>
      </el-table-column>
      <el-table-column
        fixed="right"
        label="操作"
        align="center"
        width="200"
        class-name="small-padding fixed-width"
      >
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="handleEdit(scope.row)" style="width:64px">编辑</el-button>
          <!-- <el-button size="mini" type="danger" @click="handleDelete(scope.row)" style="width:64px">解散仙盟</el-button> -->
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination-container" style="margin-top:15px;">
      <el-pagination
        :current-page="listQuery.pageIndex"
        :page-sizes="[20]"
        :total="total"
        background
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="handleCurrentChange"
      />
    </div>
    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form
        ref="dataForm"
        :model="temp"
        label-position="left"
        label-width="80px"
        style="width: 460px; margin-left:50px;"
      >
        <el-form-item v-if="dialogStatus=='update'" label="id">
          <el-input v-model="temp.id" type="number" :disabled="true" />
        </el-form-item>

        <el-form-item label="仙盟公告">
          <el-input
            v-model="temp.notice"
            :model="temp"
            autosize
            type="textarea"
            placeholder="仙盟公告"
            :disabled="dialogStatus=='update'"
          />
        </el-form-item>
      </el-form>
      <el-input v-model="temp.id" type="number" :disabled="true" />
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="danger" @click="updateData">确认修改</el-button>
      </span>
    </el-dialog>
    <el-dialog :visible.sync="dialogPvVisible" title="是否确认解散" width="30%">
      <div>是否确认解散仙盟</div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogPvVisible = false">取消</el-button>
        <el-button type="danger" @click="disMiss()">解散</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import { parseTime } from "@/utils/index";
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { Message, MessageBox } from "element-ui";
import { getAllianceList } from "@/api/alliance";
import { allianceGongGaoForm } from "@/api/alliance";
import { allianceDismissForm } from "@/api/alliance";

export default {
  name: "AllianceList",
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
        delete: "解散仙盟"
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
      this.getList();
    },

    getList() {
      this.listLoading = true;
      getAllianceList(this.listQuery)
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
    handleEdit(e) {
      this.dialogFormVisible = true;
      this.temp = Object.assign({}, e);
    },
    handleDelete(e) {
      this.dialogPvVisible = true;
      this.temp = Object.assign({}, e);
    },
    disMiss() {
      const postData ={
         allianceId:this.temp.id,
         serverId:this.listQuery.serverId
      }
      console.log(postData)
      allianceDismissForm(postData).then(res=> {
         this.dialogPvVisible = false;
         this.showSuccess(); 
        this.handleFilter();
      })
    
    },
    updateData(e) {
       const postData = {
        allianceId:this.temp.id,
        gongGao:this.temp.notice,
        serverId: this.listQuery.serverId
     };
      allianceGongGaoForm(postData).then(res => {
        this.dialogFormVisible = false;
        this.showSuccess();
        this.handleFilter();
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

