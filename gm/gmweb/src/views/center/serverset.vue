<template>
    
    <div class="app-container">
        <div class="filter-container">
            <el-select v-model="listQuery.platformId" class="filter-item" clearable placeholder="中心平台" disabled >
                <el-option v-for="citem in centerPlatformList" :key="citem.centerPlatformId" :label="citem.centerPlatformName" :value="citem.centerPlatformId"/>
            </el-select>
            <el-input v-model="listQuery.centerServerName" placeholder="中心服务名" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" disabled />
        </div>
        <el-row :gutter="20">
            <el-col :span="12">
                <div>未添加</div>
                <el-table
                    v-loading="listLoading"
                    :key="tableKey"
                    :data="unList"
                    border
                    fit
                    highlight-current-row
                    style="width: 100%;margin-top:15px;">
                    <el-table-column label="中心服务ID" align="center" width="150px">
                        <template slot-scope="scope">
                        <span>{{ scope.row.id }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column label="中心服务名" width="150px" align="left">
                        <template slot-scope="scope">
                        <span>{{ scope.row.serverName }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column label="服务ID序号" width="120px" align="left">
                        <template slot-scope="scope">
                        <span>{{ scope.row.serverId }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column label="已绑战区服务序号" width="140px" align="left">
                        <template slot-scope="scope">
                        <span>{{ scope.row.parentServerId }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column label="操作" align="center" width="150" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                            <el-button type="primary" size="mini" @click="handleUpdate(scope.row)">绑定</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-col>
            <el-col :span="12">
                <div>已添加</div>   
                <el-table
                    v-loading="listLoading"
                    :key="tableKeyEn"
                    :data="enList"
                    border
                    fit
                    highlight-current-row
                    style="width: 100%;margin-top:15px;">
                    <el-table-column label="中心服务ID" align="center" width="150px">
                        <template slot-scope="scope">
                        <span>{{ scope.row.id }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column label="中心服务名" width="150px" align="left">
                        <template slot-scope="scope">
                        <span>{{ scope.row.serverName }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column label="服务ID序号" width="150px" align="left">
                        <template slot-scope="scope">
                        <span>{{ scope.row.serverId }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column label="操作" align="center" width="150" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                            <el-button type="danger" size="mini" @click="handleDelete(scope.row)">解绑</el-button>
                        </template>
                    </el-table-column>
                </el-table>     
            </el-col>
        </el-row>
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
  updateCenterServerParentId
} from "@/api/centerserver";
import { getCenterPlatList } from "@/api/center";

export default {
  name: "serverSet",
  directives: {
    waves
  },
  data() {
    return {
      listQuery: {
        pageIndex: 1,
        centerServerName: "",
        platformId: undefined,
        serverType: 0
      },
      listLoading: false,
      tableKey: "serverset",
      tableKeyEn: "serverseten",
      unList: [],
      enList: [],
      platformId: undefined,
      serverId: undefined,
      centerPlatformList: [],
      zhanChangServerList: []
    };
  },
  created() {
    this.initData();
  },
  methods: {
    initData() {
      const platFormId = this.$route.params && this.$route.params.platformId;
      const serverId = this.$route.params && this.$route.params.serverId;
      const serverName = this.$route.params && this.$route.params.serverName;
      this.serverId = serverId;
      this.platformId = platFormId;
      this.listQuery.platformId = platFormId;
      this.listQuery.centerServerName = serverName;

      getCenterPlatList().then(res => {
        this.centerPlatformList = res.itemArray;
      });
      this.loadData();
    },
    loadData() {
      let that = this;
      getCenterServerListByServerType(this.platformId, 0).then(res => {
        let itemArray = res.itemArray;
        that.unList = [];
        that.enList = [];
        if (itemArray.length > 0) {
          for (let i = 0; i < itemArray.length; i++) {
            let item = itemArray[i];
            if (item.parentServerId === that.serverId) {
              that.enList.push(item);
              continue;
            }
            that.unList.push(item);
          }
        }
      });
    },
    handleUpdate(e) {
      let id = e.id;
      let parentServerId = this.serverId;
      updateCenterServerParentId(id, parentServerId).then(res => {
        this.showSuccess();
        this.loadData();
      });
    },
    handleDelete(e) {
      let id = e.id;
      let parentServerId = 0;
      updateCenterServerParentId(id, parentServerId).then(res => {
        this.showSuccess();
        this.loadData();
      });
    },
    showSuccess() {
      this.$message({
        message: "修改成功",
        type: "success",
        duration: 1000
      });
    }
  }
};
</script>
