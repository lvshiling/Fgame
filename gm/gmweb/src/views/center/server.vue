<template>
  <div class="app-container">
    <div class="filter-container">
      <el-select v-model="listQuery.platformId" class="filter-item" clearable placeholder="中心平台">
         <el-option v-for="citem in centerPlatformList" :key="citem.centerPlatformId" :label="citem.centerPlatformName" :value="citem.centerPlatformId"/>
      </el-select>
      <el-select v-model="listQuery.serverType" class="filter-item" clearable placeholder="服务类型">
          <el-option v-for="item in serverTypeArray" :key="item.key" :label="item.name" :value="item.key"/>
      </el-select>
      <el-input v-model="listQuery.centerServerName" placeholder="中心服务名" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
      <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">添加</el-button>
    </div>

    <el-table
      v-loading="listLoading"
      :key="tableKey"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;margin-top:15px;">
      <el-table-column fixed="left" label="中心服务ID" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="left" label="中心服务名" width="150px" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverName }}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="left" label="中心服务类型" width="150px" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverType | serverTypeFilter }}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="left" label="服务ID序号" width="150px" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverId }}</span>
        </template>
      </el-table-column>
      <el-table-column label="中心平台名" width="150px" align="left">
        <template slot-scope="scope">
          <span>{{ getPlatformName(scope.row.platformId) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="服务器ip" width="120px" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverIp }}</span>
        </template>
      </el-table-column>
      <el-table-column label="服务器端口" width="100" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverPort }}</span>
        </template>
      </el-table-column>
      <el-table-column label="远程服务器ip" width="120px" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverRemoteIp }}</span>
        </template>
      </el-table-column>
      <el-table-column label="远程服务器端口" width="120" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverRemotePort }}</span>
        </template>
      </el-table-column>
      <el-table-column label="数据库IP" width="120px" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverDBIp }}</span>
        </template>
      </el-table-column>
      <el-table-column label="数据库端口" width="120" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverDBPort }}</span>
        </template>
      </el-table-column>
      <el-table-column label="数据库名" width="120" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverDBName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="数据库用户名" width="120" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverDBUser }}</span>
        </template>
      </el-table-column>
       <el-table-column label="数据库密码" width="120" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverDBPassword }}</span>
        </template>
      </el-table-column>
      <el-table-column label="服务器标签" width="120" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverTag | serverTagFilter }}</span>
        </template>
      </el-table-column>
      <el-table-column label="服务器状态" width="100" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverStatus | serverStatusFilter }}</span>
        </template>
      </el-table-column>
      <el-table-column label="开始时间" min-width="160px" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.startTime | parseTimeFilter }}</span>
        </template>
      </el-table-column>
      <el-table-column label="提前展示" min-width="160px" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.preShow | parseYesOrNo }}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" align="center" width="240" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" v-if="scope.row.serverType === 0" @click="handlePing(scope.row)">ping</el-button>
          <el-button type="primary" size="mini" v-if="scope.row.serverType === 2" @click="handleView(scope.row)">配置</el-button>
          <el-button type="primary" size="mini" @click="handleUpdate(scope.row)">编辑</el-button>
          <el-button size="mini" type="danger" @click="handleDelete(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <div class="pagination-container" style="margin-top:15px;">
      <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"/>
    </div>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :model="temp" label-position="left" label-width="100px" style="width: 400px; margin-left:50px;">
        <el-form-item label="中心服务名">
          <el-input v-model="temp.serverName"/>
        </el-form-item>

        <el-form-item label="服务类型">
          <el-select v-model="temp.serverType" class="filter-item" placeholder="服务类型" @change="handleServerTypeChange">
            <el-option v-for="item in serverTypeArray" :key="item.key" :label="item.name" :value="item.key"/>
          </el-select>
        </el-form-item>

        <el-form-item label="服务ID">
          <el-input v-model="temp.serverId" type="number"/>
        </el-form-item>

        <el-form-item v-if="temp.serverType !== 4 && temp.serverType != 5" label="中心平台">
          <el-select v-model="temp.platformId" class="filter-item" placeholder="中心平台" @change="handleCenterPlatformTypeChange">
            <el-option v-for="citem in centerPlatformList" :key="citem.centerPlatformId" :label="citem.centerPlatformName" :value="citem.centerPlatformId"/>
          </el-select>
        </el-form-item>

        <el-form-item v-if="temp.serverType === 0" label="归属战区">
          <el-select v-model="temp.parentServerId" class="filter-item" placeholder="战区服务器">
            <el-option v-for="citem in zhanChangServerList" :key="citem.serverId" :label="citem.serverName" :value="citem.serverId"/>
          </el-select>
        </el-form-item>

        <el-form-item v-if="temp.serverType === 0"  label="服务IP">
          <el-input v-model="temp.serverIp" />
        </el-form-item>
        <el-form-item v-if="temp.serverType === 0" label="服务端口">
          <el-input v-model="temp.serverPort" type="number"/>
        </el-form-item>

        <el-form-item v-if="temp.serverType === 0" label="远程服务器ip">
          <el-input v-model="temp.serverRemoteIp" />
        </el-form-item>
        <el-form-item v-if="temp.serverType === 0" label="远程服务器端口">
          <el-input v-model="temp.serverRemotePort" type="number"/>
        </el-form-item>

        <el-form-item v-if="temp.serverType === 0" label="数据库IP">
          <el-input v-model="temp.serverDBIp" />
        </el-form-item>
        <el-form-item v-if="temp.serverType === 0" label="数据库端口">
          <el-input v-model="temp.serverDBPort" type="number"/>
        </el-form-item>
        <el-form-item v-if="temp.serverType === 0" label="数据库名字">
          <el-input v-model="temp.serverDBName" />
        </el-form-item>
        <el-form-item v-if="temp.serverType === 0" label="数据库用户名">
          <el-input v-model="temp.serverDBUser" />
        </el-form-item>
        <el-form-item v-if="temp.serverType === 0" label="数据库密码">
          <el-input v-model="temp.serverDBPassword" />
        </el-form-item>

        <el-form-item v-if="temp.serverType === 0" label="服务器标签">
          <el-select v-model="temp.serverTag" class="filter-item" placeholder="服务器标签">
            <el-option v-for="citem in serverTagArray" :key="citem.key" :label="citem.name" :value="citem.key"/>
          </el-select>
        </el-form-item>

        <el-form-item v-if="temp.serverType === 0" label="服务器状态">
          <el-select v-model="temp.serverStatus" class="filter-item" placeholder="服务器状态">
            <el-option v-for="citem in serverStatusArray" :key="citem.key" :label="citem.name" :value="citem.key"/>
          </el-select>
        </el-form-item>

        <el-form-item v-if="temp.serverType === 0" label="开始时间">
          <el-date-picker v-model="temp.startTimestr" type="datetime" placeholder="选择开始日期时间"/>
        </el-form-item>

        <el-form-item v-if="temp.serverType === 0" label="提前展示">
          <el-checkbox v-model="temp.preShow">开服前30分钟展示</el-checkbox>
        </el-form-item>

      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button v-if="dialogStatus=='create'" type="primary" @click="createData">创建</el-button>
        <el-button v-else type="primary" @click="updateData">确定</el-button>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="dialogPvVisible" title="是否确认删除" width="30%">
      <div>
        是否确认删除中心服务：{{ temp.serverName }}
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogPvVisible = false">取消</el-button>
        <el-button type="danger" @click="deleteData">删除</el-button>
      </span>
    </el-dialog>

    <el-dialog :visible.sync="dialogPingFormVisible" title="是否确认Ping服务器" width="30%">
      <div>
        是否确认Ping中心服务：{{ temp.serverName }}
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogPingFormVisible = false">取消</el-button>
        <el-button type="danger" @click="pingData">ping</el-button>
      </span>
    </el-dialog>

  

  </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import {
  getCenterServerList,
  getCenterServerAdd,
  getCenterServerUpdate,
  getCenterServerDelete,
  refreshCenterServer,
  getCenterServerListByServerType,
  updateCenterServerParentId,
  getCenterServerPing
} from "@/api/centerserver";
import { getCenterPlatList } from "@/api/center";
import {
  serverTypeList,
  serverTagList,
  serverStatusList
} from "@/types/center";
import { parseTime } from "@/utils/index";

export default {
  name: "CenterServerList",
  directives: {
    waves
  },
  filters: {
    parseTimeFilter: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
    },
    parseYesOrNo: function(value) {
      if (value) {
        return "是";
      }
      return "否";
    },
    serverTypeFilter: function(value) {
      return serverTypeList[value].name;
    },
    serverTagFilter: function(value) {
      return serverTagList[value].name;
    },
    serverStatusFilter: function(value) {
      return serverStatusList[value].name;
    }
  },
  data() {
    return {
      listLoading: false,
      tableKey: 0,
      total: 0,
      listQuery: {
        pageIndex: 1,
        centerServerName: "",
        platformId: undefined,
        serverType: undefined
      },
      textMap: {
        update: "编辑",
        create: "添加",
        refresh: "刷新"
      },
      serverTypeArray: [],
      serverTagArray: [],
      serverStatusArray: [],
      centerPlatformList: [],
      zhanChangServerList: [],

      dialogStatus: "",
      dialogPvVisible: false,
      dialogFormVisible: false,
      dialogPingFormVisible:false,
      temp: {
        id: 0,
        serverType: undefined,
        serverId: undefined,
        platformId: undefined,
        serverName: undefined,
        startTimestr: new Date(),
        startTime: undefined
      },
      list: []
    };
  },
  created() {
    this.initData();
    this.getList();
  },
  methods: {
    handleFilter: function() {
      this.listQuery.pageIndex = 1;
      this.getList();
    },
    handleCreate: function() {
      this.dialogStatus = "create";
      this.dialogFormVisible = true;
      this.temp = {
        id: 0,
        serverType: undefined,
        serverId: undefined,
        platformId: undefined,
        serverName: undefined,
        startTimestr: new Date(),
        startTime: undefined,
        serverIp: undefined,
        serverPort: undefined,
        serverRemoteIp: undefined,
        serverRemotePort: undefined,
        serverDBIp: undefined,
        serverDBPort: undefined,
        serverDBName: undefined,
        serverDBUser: undefined,
        serverDBPassword: undefined,
        serverTag: undefined,
        serverStatus: undefined,
        parentServerId: undefined,
        preShow: false
      };
    },
    handleUpdate: function(e) {
      const curRow = Object.assign({}, e);
      //   this.temp.startTimestr = new Date(this.temp.startTime);
      this.temp = {
        id: curRow.id,
        serverType: curRow.serverType,
        serverId: curRow.serverId,
        platformId: curRow.platformId,
        serverName: curRow.serverName,
        startTimestr: new Date(curRow.startTime),
        startTime: curRow.startTime,
        serverIp: curRow.serverIp,
        serverPort: curRow.serverPort,
        serverRemoteIp: curRow.serverRemoteIp,
        serverRemotePort: curRow.serverRemotePort,
        serverDBIp: curRow.serverDBIp,
        serverDBPort: curRow.serverDBPort,
        serverDBName: curRow.serverDBName,
        serverDBUser: curRow.serverDBUser,
        serverDBPassword: curRow.serverDBPassword,
        serverTag: curRow.serverTag,
        serverStatus: curRow.serverStatus,
        parentServerId: curRow.parentServerId,
        preShow: curRow.preShow
      };
      getCenterServerListByServerType(curRow.platformId, 2).then(res => {
        this.zhanChangServerList = res.itemArray;
      });
      this.dialogStatus = "update";
      this.dialogFormVisible = true;
    },
    handlePing: function(e) {
      const curRow = Object.assign({}, e);
      //   this.temp.startTimestr = new Date(this.temp.startTime);
      this.temp = {
        id: curRow.id,
        serverId: curRow.serverId,
        platformId: curRow.platformId
      };
      this.dialogStatus = "ping";
      this.dialogPingFormVisible = true;
    },
    handleView(e) {
      const curRow = Object.assign({}, e);
      let serverId = curRow.serverId;
      let platFormId = curRow.platformId;
      let serverName = curRow.serverName;
      console.log("push");
      this.$router.push({
        name: "serverset",
        params: {
          serverId: serverId,
          platformId: platFormId,
          serverName: serverName
        }
      });
    },
    handleDelete: function(e) {
      this.dialogPvVisible = true;
      this.temp = Object.assign({}, e);
    },
    handleServerTypeChange: function(e) {
      if (e === 4) {
        this.temp.platformId = undefined;
      }
      if (this.temp.platformId == 0) {
        this.temp.platformId = undefined;
      }
      if (e !== 0) {
        this.temp.startTime = undefined;
        this.temp.serverIp = undefined;
        this.temp.serverPort = undefined;
        this.temp.serverRemoteIp = undefined;
        this.temp.serverRemotePort = undefined;
        this.temp.serverDBIp = undefined;
        this.temp.serverDBPort = undefined;
        this.temp.serverDBName = undefined;
        this.temp.serverDBUser = undefined;
        this.temp.serverDBPassword = undefined;
        this.temp.serverTag = undefined;
        this.temp.serverStatus = undefined;
      }
    },
    handleCenterPlatformTypeChange: function(e) {
      getCenterServerListByServerType(this.temp.platformId, 2).then(res => {
        this.zhanChangServerList = res.itemArray;
      });
    },
    getList() {
      this.listLoading = true;
      getCenterServerList(this.listQuery)
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
    updateData() {
      this.temp.startTime = this.temp.startTimestr.valueOf();
      this.temp.serverId = parseInt(this.temp.serverId);
      this.temp.serverTag = parseInt(this.temp.serverTag);
      this.temp.serverStatus = parseInt(this.temp.serverStatus);
      getCenterServerUpdate(this.temp).then(() => {
        this.getList();
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    createData() {
      this.temp.startTime = this.temp.startTimestr.valueOf();
      this.temp.serverId = parseInt(this.temp.serverId);
      this.temp.serverTag = parseInt(this.temp.serverTag);
      this.temp.serverStatus = parseInt(this.temp.serverStatus);
      getCenterServerAdd(this.temp).then(() => {
        this.getList();
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    deleteData() {
      getCenterServerDelete(this.temp).then(() => {
        this.getList();
        this.dialogPvVisible = false;
        this.showSuccess();
      });
    },
    pingData(){
      getCenterServerPing(this.temp).then(() => {
        this.getList();
        this.dialogPingFormVisible = false;
        this.showSuccessMessage("ping成功");
      });
    },
    getPlatformName(value) {
      for (let i = 0, len = this.centerPlatformList.length; i < len; i++) {
        const item = this.centerPlatformList[i];
        if (item.centerPlatformId == value) {
          return item.centerPlatformName;
        }
      }
    },
    initData() {
      this.serverTypeArray = serverTypeList;
      this.serverTagArray = serverTagList;
      this.serverStatusArray = serverStatusList;
      getCenterPlatList().then(res => {
        this.centerPlatformList = res.itemArray;
      });
    },
    showSuccess() {
      this.$message({
        message: "修改成功",
        type: "success",
        duration: 1000
      });
    },
    showSuccessMessage(msg) {
      this.$message({
        message: msg,
        type: "success",
        duration: 1000
      });
    }
  }
};
</script>

