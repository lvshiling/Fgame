<template>
    <div class="app-container">
        <div class="filter-container">
            <el-input placeholder="平台名称" v-model="listQuery.platformName" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
            <el-select v-model="listQuery.channelId" placeholder="渠道" clearable style="width: 120px" class="filter-item">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
            <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">添加</el-button>
            <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleRefreshSdk">刷新Sdk</el-button>
        </div>

        <el-table
            v-loading="listLoading"
            :key="tableKey"
            :data="list"
            border
            fit
            highlight-current-row
            style="width: 100%;margin-top:15px;">
            <el-table-column label="平台ID" align="center" width="65px">
                <template slot-scope="scope">
                    <span>{{ scope.row.platformId }}</span>
                </template>
            </el-table-column>
             <el-table-column label="平台名" min-width="150px" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.platformName}}</span>
                </template>
            </el-table-column>
            <el-table-column label="渠道" min-width="150px" align="center">
                <template slot-scope="scope">
                    <span>{{ channelName(scope.row.channelId)}}</span>
                </template>
            </el-table-column>
             <el-table-column label="Sdk类型" min-width="150px" align="center">
                <template slot-scope="scope">
                    <span>{{ getSdkTypeName(scope.row.sdkType) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作" align="center" width="200" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                <el-button type="primary" size="mini" @click="handleUpdate(scope.row)">编辑</el-button>
                <el-button size="mini" type="danger" @click="handleDelete(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>

        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>

        <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
        <el-form ref="dataForm" :model="temp" label-position="left" label-width="70px" style="width: 500px; margin-left:50px;">
            <el-form-item label="平台名">
                <el-input v-model="temp.platformName"/>
            </el-form-item>
            <el-form-item label="渠道">
                <el-select v-model="temp.channelId" class="filter-item" placeholder="渠道">
                     <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
                </el-select>
            </el-form-item>

            <el-form-item label="中心平台">
                <el-select v-model="temp.centerPlatformId" class="filter-item" placeholder="中心平台">
                     <el-option v-for="citem in centerPlatformList" :key="citem.centerPlatformId" :label="citem.centerPlatformName" :value="citem.centerPlatformId"/>
                </el-select>
            </el-form-item>

            <el-form-item label="sdk类型">
              <el-select v-model="temp.sdkType" placeholder="sdk类型" style="width: 160px" class="filter-item" >
                <el-option v-for="item in skdTypeArray" :key="item.key" :label="item.name" :value="item.key" />
              </el-select>
            </el-form-item>
            <el-form-item label="api密钥">
              <el-row>
                <el-col :span="20"><el-input v-model="temp.signKey"/></el-col>
                <el-col :span="4"><el-button type="primary" size="mini" class="row-left" @click="createUUID">生成</el-button></el-col>
              </el-row>
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
              是否确认删除平台：{{temp.platformName}}
          </div>
          <span slot="footer" class="dialog-footer">
            <el-button @click="dialogPvVisible = false">取消</el-button>
            <el-button type="danger" @click="deleteData">删除</el-button>
          </span>
        </el-dialog>

        <el-dialog :visible.sync="dialogRefreshSdkVisible" title="是否确认刷新" width="30%">
          <div>
              是否确认刷新
          </div>
          <span slot="footer" class="dialog-footer">
            <el-button @click="dialogRefreshSdkVisible = false">取消</el-button>
            <el-button type="primary" @click="refreshSdk">刷新</el-button>
          </span>
        </el-dialog>

        
    </div>
</template>
<style>
.row-left {
  margin-left: 10px;
}
</style>


<script>
import waves from "@/directive/waves"; // 水波纹指令
import {
  getPlatformList,
  setPlatformAdd,
  setPlatformUpdate,
  setPlatformDelete,
  refreshSdk
} from "@/api/platform";
import { getAllChannel } from "@/api/channel";
import { getCenterPlatList } from "@/api/center";
import { sdkTypeList } from "@/types/center";
import { getAllSdkType } from "@/api/center";
import { newUUID } from "@/utils/tool";
export default {
  name: "PlatformList",
  directives: {
    waves
  },
  filters: {
    parseSdkType: function(value) {
      if (sdkTypeList[value - 1]) {
        return sdkTypeList[value - 1].name;
      }
      return "";
    }
  },
  created() {
    this.skdTypeArray = sdkTypeList;
    this.getAllChannel();
    this.getList();
    this.initCenter();
    getAllSdkType().then(res => {
      this.skdTypeArray = res.itemArray;
    });
  },
  data() {
    return {
      listLoading: false,
      tableKey: 1,
      total: 0,
      listQuery: {
        pageIndex: 1,
        platformName: "",
        channelId: undefined,
        platformName: undefined
      },
      textMap: {
        update: "编辑",
        create: "添加"
      },
      dialogStatus: "",
      dialogPvVisible: false,
      dialogFormVisible: false,
      dialogRefreshSdkVisible: false,
      channelList: [],
      temp: {},
      list: [],
      centerPlatformList: [],
      skdTypeArray: []
    };
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
        centerPlatformId: undefined,
        channelId: undefined,
        sdkType: undefined,
        signKey:""
      };
    },
    handleRefreshSdk: function() {
      this.dialogRefreshSdkVisible = true;
    },
    handleUpdate: function(e) {
      this.dialogStatus = "update";
      this.dialogFormVisible = true;
      this.temp = Object.assign({}, e);
      if (!this.temp.centerPlatformId) {
        this.temp.centerPlatformId = undefined;
      }
    },
    handleDelete: function(e) {
      this.dialogPvVisible = true;
      this.temp = Object.assign({}, e);
    },
    createUUID: function(e) {
      this.temp.signKey = newUUID();
    },
    getAllChannel() {
      getAllChannel().then(res => {
        this.channelList = res.itemArray;
      });
    },
    getList() {
      this.listLoading = true;
      getPlatformList(this.listQuery)
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
      setPlatformUpdate(this.temp).then(() => {
        this.getList();
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    createData() {
      setPlatformAdd(this.temp).then(() => {
        this.getList();
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    deleteData() {
      setPlatformDelete(this.temp).then(() => {
        this.getList();
        this.dialogPvVisible = false;
        this.showSuccess();
      });
    },
    refreshSdk() {
      refreshSdk().then(() => {
        this.dialogRefreshSdkVisible = false;
      });
    },
    showSuccess() {
      this.$message({
        message: "修改成功",
        type: "success",
        duration: 1000
      });
    },
    channelName: function(channelId) {
      let channelName = this.channelList.find(n => {
        return n.channelId == channelId;
      });
      if (channelName && channelName.channelName) {
        return channelName.channelName;
      }
      return "";
    },
    initCenter() {
      getCenterPlatList().then(res => {
        this.centerPlatformList = res.itemArray;
      });
    },
    getSdkTypeName(value) {
      for (let i = 0, len = this.skdTypeArray.length; i < len; i++) {
        if (this.skdTypeArray[i].key == value) {
          return this.skdTypeArray[i].name;
        }
      }
      return value;
    }
  }
};
</script>

