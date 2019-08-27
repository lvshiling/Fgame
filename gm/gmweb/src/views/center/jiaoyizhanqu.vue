<template>
  <div class="app-container">
    <div class="filter-container">
      <el-select v-model="listQuery.platformId" class="filter-item" clearable placeholder="中心平台">
         <el-option v-for="citem in centerPlatformList" :key="citem.centerPlatformId" :label="citem.centerPlatformName" :value="citem.centerPlatformId"/>
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
      <el-table-column fixed="left" label="交易战区主键ID" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="中心平台名" width="150px" align="left">
        <template slot-scope="scope">
          <span>{{ getPlatformName(scope.row.platformId) }}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="left" label="交易战区序号" width="150px" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.serverId }}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="left" label="交易战区名" width="150px" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.zhanquName }}</span>
        </template>
      </el-table-column>

      <el-table-column label="创建时间" min-width="160px" align="left">
        <template slot-scope="scope">
          <span>{{ scope.row.createTime | parseTimeFilter }}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" align="center" width="240" class-name="small-padding fixed-width">
        <template slot-scope="scope">
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
        <el-form-item label="交易战区服务名">
          <el-input v-model="temp.zhanquName"/>
        </el-form-item>

        <el-form-item label="交易战区序号">
          <el-input v-model="temp.serverId" type="number"/>
        </el-form-item>

        <el-form-item label="中心平台">
          <el-select v-model="temp.platformId" class="filter-item" placeholder="中心平台">
            <el-option v-for="citem in centerPlatformList" :key="citem.centerPlatformId" :label="citem.centerPlatformName" :value="citem.centerPlatformId"/>
          </el-select>
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
        是否确认删除中心交易战区：{{ temp.serverName }}
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogPvVisible = false">取消</el-button>
        <el-button type="danger" @click="deleteData">删除</el-button>
      </span>
    </el-dialog>

  

  </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import {
  addJiaoYiZhanQu,
  updateJiaoYiZhanQu,
  deleteJiaoYiZhanQu,
  getJiaoYiZhanQuList
} from "@/api/jiaoyizhanqu";
import { getCenterPlatList } from "@/api/center";
import { parseTime } from "@/utils/index";

export default {
  name: "CenterJiaoYiZhanQuList",
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
      centerPlatformList:[],
      dialogStatus: "",
      dialogPvVisible: false,
      dialogFormVisible: false,
      temp: {
        id: 0,
        serverId: undefined,
        platformId: undefined,
        zhanquName: undefined,
        createTime: undefined
      },
      list: []
    };
  },
  created() {
    this.initData();
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
        serverId: undefined,
        platformId: undefined,
        zhanquName: undefined,
        createTime: undefined
      };
    },
    handleUpdate: function(e) {
      const curRow = Object.assign({}, e);
      //   this.temp.startTimestr = new Date(this.temp.startTime);
      this.temp = {
        id: curRow.id,
        serverId: curRow.serverId,
        platformId: curRow.platformId,
        zhanquName: curRow.zhanquName,
        createTime: curRow.createTime
      };
      this.dialogStatus = "update";
      this.dialogFormVisible = true;
    },
    handleDelete: function(e) {
      this.dialogPvVisible = true;
      this.temp = Object.assign({}, e);
    },
    
    handleCenterPlatformTypeChange: function(e) {
      getCenterServerListByServerType(this.temp.platformId, 2).then(res => {
        this.zhanChangServerList = res.itemArray;
      });
    },
    getList() {
      if(!this.listQuery.platformId){
        this.showError("请选择平台");
        return;
      }
      this.listLoading = true;
      getJiaoYiZhanQuList(this.listQuery)
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
      updateJiaoYiZhanQu(this.temp).then(() => {
        this.getList();
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    createData() {
      addJiaoYiZhanQu(this.temp).then(() => {
        this.getList();
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    deleteData() {
      deleteJiaoYiZhanQu(this.temp).then(() => {
        this.getList();
        this.dialogPvVisible = false;
        this.showSuccess();
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
    showError(msg){
        this.$message({
        message: msg,
        type: "error",
        duration: 1000
      });
    }
  }
};
</script>

