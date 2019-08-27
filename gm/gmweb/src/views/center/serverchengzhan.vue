<template>
    <div class="app-container">
        <div class="filter-container">
            <!-- <el-select v-model="listQuery.platformId" class="filter-item" placeholder="中心平台" >
                <el-option v-for="citem in centerPlatformList" :key="citem.centerPlatformId" :label="citem.centerPlatformName" :value="citem.centerPlatformId"/>
            </el-select>
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button> -->
            城战平台服配置
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleBindMultiple">批量绑定</el-button>
        </div>
        <el-row :gutter="20">
            <el-col :span="12">
                <div>未添加服务</div>
                <el-table
                    v-loading="listLoading"
                    :key="tableKey"
                    :data="unList"
                    border
                    fit
                    highlight-current-row
                    style="width: 100%;margin-top:15px;"
                    @selection-change="handleSelectionChange">
                    <el-table-column
                      type="selection"
                      width="55">
                    </el-table-column>
                    <el-table-column label="主键ID" align="center" width="80px">
                        <template slot-scope="scope">
                        <span>{{ scope.row.id }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column label="中心平台名" width="120px" align="left">
                      <template slot-scope="scope">
                        <span>{{ getPlatformName(scope.row.platformId) }}</span>
                      </template>
                    </el-table-column>
                    <el-table-column label="中心服务名" width="150px" align="left">
                        <template slot-scope="scope">
                        <span>{{ scope.row.serverName }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column label="服务ID序号" width="80px" align="left">
                        <template slot-scope="scope">
                        <span>{{ scope.row.serverId }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column label="合服服务ID" width="80px" align="left">
                        <template slot-scope="scope">
                        <span>{{ scope.row.finnalServerId }}</span>
                        </template>
                    </el-table-column>
                    <!-- <el-table-column label="已绑交易战区序号" width="140px" align="left">
                        <template slot-scope="scope">
                        <span>{{ scope.row.chengZhanServerId }}</span>
                        </template>
                    </el-table-column> -->
                    <el-table-column label="操作" align="center" width="150" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                            <el-button type="primary" size="mini" @click="handleBind(scope.row)">绑定</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-col>
            <el-col :span="12">
                <div v-for="(item,index) in quanPingTaiServerList" :key="index">
                <div>已添加,城战平台服务ID:{{item.serverId}},名称:{{item.serverName}}</div>   
                    <el-table
                        v-loading="listLoading"
                        :key="tableKeyEn"
                        :data="jiaoYiZhanQuServerMap[item.serverId]"
                        border
                        fit
                        highlight-current-row
                        style="width: 100%;margin-top:15px;">
                        <el-table-column label="主键ID" align="center" width="80px">
                            <template slot-scope="scope">
                            <span>{{ scope.row.id }}</span>
                            </template>
                        </el-table-column>
                        <el-table-column label="中心平台名" width="120px" align="left">
                          <template slot-scope="scope">
                            <span>{{ getPlatformName(scope.row.platformId) }}</span>
                          </template>
                        </el-table-column>
                        <el-table-column label="中心服务名" width="150px" align="left">
                            <template slot-scope="scope">
                            <span>{{ scope.row.serverName }}</span>
                            </template>
                        </el-table-column>
                        <el-table-column label="服务ID序号" width="80px" align="left">
                            <template slot-scope="scope">
                            <span>{{ scope.row.serverId }}</span>
                            </template>
                        </el-table-column>
                        <el-table-column label="合服服务ID" width="80px" align="left">
                            <template slot-scope="scope">
                            <span>{{ scope.row.finnalServerId }}</span>
                            </template>
                        </el-table-column>
                        <el-table-column label="操作" align="center" width="150" class-name="small-padding fixed-width">
                            <template slot-scope="scope">
                                <el-button type="danger" size="mini" @click="handleDelete(scope.row)">解绑</el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                    <br />
                </div>     
            </el-col>
        </el-row>

        <el-dialog title="绑定城战平台" :visible.sync="dialogFormVisible">
            <el-form ref="dataForm" :model="temp" label-position="left" label-width="100px" style="width: 400px; margin-left:50px;">
                <el-form-item label="服务器名">
                <el-input v-model="temp.serverName" disabled />
                </el-form-item>

                <el-form-item label="服务器序号">
                <el-input v-model="temp.serverId" type="number"  disabled />
                </el-form-item>

                <el-form-item label="城战平台">
                    <el-select v-model="temp.chengZhanServerId" class="filter-item" placeholder="城战平台">
                        <el-option v-for="citem in quanPingTaiServerList" :key="citem.serverId" :label="citem.serverId" :value="citem.serverId"/>
                    </el-select>
                </el-form-item>
            </el-form>

            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogFormVisible = false">取消</el-button>
                <el-button type="primary" @click="updateBind">确定</el-button>
            </div>
        </el-dialog>

        <el-dialog title="批量绑定城战平台" :visible.sync="dialogMupliFormVisible">
            <el-form ref="dataForm" :model="temp" label-position="left" label-width="100px" style="width: 400px; margin-left:50px;">
                <el-form-item label="城战平台">
                    <el-select v-model="temp.chengZhanServerId" class="filter-item" placeholder="城战平台">
                        <el-option v-for="citem in quanPingTaiServerList" :key="citem.serverId" :label="citem.serverId" :value="citem.serverId"/>
                    </el-select>
                </el-form-item>
            </el-form>

            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogMupliFormVisible = false">取消</el-button>
                <el-button type="primary" @click="updateBindArray">确定</el-button>
            </div>
        </el-dialog>

        <el-dialog :visible.sync="dialogPvVisible" title="是否确认删除绑定" width="30%">
            <div>
                是否确认解绑城战平台：{{ temp.serverName }}
            </div>
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogPvVisible = false">取消</el-button>
                <el-button type="danger" @click="deleteData">解绑</el-button>
            </span>
        </el-dialog>
    </div>        
        
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import {
  getCenterServerListByServerType,
  updateCenterServerChengZhan,
  updateCenterServerChengZhanArray,
  getCenterServerListChengZhan
} from "@/api/centerserver";

import { getCenterPlatList } from "@/api/center";

export default {
  name: "serverJiaoYiZhanQu",
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
      tableKey: "serverjiaoyi",
      tableKeyEn: "serverjiaoyien",
      unList: [],
      jiaoYiZhanQuServerMap: [],
      allServer: [],
      platformId: undefined,
      centerPlatformList: [],
      quanPingTaiServerList: [],
      dialogFormVisible: false,
      dialogPvVisible: false,
      dialogMupliFormVisible:false,
      temp: {
        id: undefined,
        serverName: undefined,
        platformId: undefined,
        serverId: undefined,
        chengZhanServerId: undefined
      },
      multipleSelection: []
    };
  },
  created() {
    this.initData();
    this.loadData();
  },
  methods: {
    initData() {
      getCenterPlatList().then(res => {
        this.centerPlatformList = res.itemArray;
      });
    },
    loadData() {
      let that = this;
      getCenterServerListChengZhan().then(res => {
        this.quanPingTaiServerList = res.itemArray;
        console.log(this.quanPingTaiServerList);
        getCenterServerListByServerType(this.listQuery, 0).then(res => {
          let itemArray = res.itemArray;
          console.log(itemArray);
          that.unList = [];
          that.jiaoYiZhanQuServerMap = [];
          if (itemArray.length > 0) {
            for (let i = 0; i < itemArray.length; i++) {
              let item = itemArray[i];
              if (item.chengZhanServerId == 0) {
                that.unList.push(item);
                continue;
              }
              for (let j = 0; j < that.quanPingTaiServerList.length; j++) {
                let jiaoYiItem = that.quanPingTaiServerList[j];
                if (jiaoYiItem.serverId == item.chengZhanServerId) {
                  if (!that.jiaoYiZhanQuServerMap[jiaoYiItem.serverId]) {
                    that.jiaoYiZhanQuServerMap[jiaoYiItem.serverId] = [];
                  }
                  that.jiaoYiZhanQuServerMap[jiaoYiItem.serverId].push(item);
                  break;
                }
              }
            }
            this.multipleSelection = [];
          }
        });
      });
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    handleFilter: function() {
      console.log(this.listQuery);
      this.listQuery.pageIndex = 1;
      this.loadData();
    },
    handleBindMultiple:function(){
      this.dialogMupliFormVisible = true;
      this.temp = {
        chengZhanServerId: undefined
      };
    },
    handleBind: function(e) {
      this.dialogFormVisible = true;
      const curRow = Object.assign({}, e);
      this.temp = {
        id: e.id,
        serverName: e.serverName,
        platformId: e.platformId,
        serverId: e.serverId,
        chengZhanServerId: undefined
      };
    },
    updateBind(e) {
      let id = this.temp.id;
      let parentServerId = this.temp.chengZhanServerId;
      updateCenterServerChengZhan(id, parentServerId).then(res => {
        this.dialogFormVisible = false;
        this.showSuccess();
        this.loadData();
      });
    },
    updateBindArray(e){
      let id = [];
      let parentServerId = this.temp.chengZhanServerId;
      if(this.multipleSelection.length == 0){
        this.showError("请选择服务器");
        return
      }
      for(let i = 0; i< this.multipleSelection.length;i++){
        let item = this.multipleSelection[i];
        id.push(item.id)
      }
      updateCenterServerChengZhanArray(id, parentServerId).then(res => {
        this.dialogMupliFormVisible = false;
        this.showSuccess();
        this.loadData();
      });
    },
    handleDelete: function(e) {
      this.dialogPvVisible = true;
      const curRow = Object.assign({}, e);
      this.temp = {
        id: e.id,
        serverName: e.serverName,
        platformId: e.platformId,
        serverId: e.serverId
      };
    },
    deleteData(e) {
      let id = this.temp.id;
      updateCenterServerChengZhan(id, 0).then(res => {
        this.showSuccess();
        this.loadData();
        this.dialogPvVisible = false;
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
    showSuccess() {
      this.$message({
        message: "修改成功",
        type: "success",
        duration: 1000
      });
    },
    showError(msg) {
      this.$message({
        message: msg,
        type: "error",
        duration: 1000
      });
    }
  }
};
</script>
