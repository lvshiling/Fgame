<template>
    <div class="app-container">
        <div class="filter-container">
            <el-input placeholder="渠道名" v-model="listQuery.channelName" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
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
            <el-table-column label="渠道ID" align="center" width="65">
                <template slot-scope="scope">
                    <span>{{ scope.row.channelId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="渠道名" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.channelName}}</span>
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
        <el-form ref="dataForm" :model="temp" label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">
            <el-form-item label="渠道名">
                <el-input v-model="temp.channelName"/>
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
              是否确认删除渠道：{{temp.channelName}}
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
import { getChannelList, getChannelAdd,getChannelUpdate,getChannelDelete } from "@/api/channel";
export default {
  name: "ChannelList",
  directives: {
    waves
  },
  created() {
    this.getList();
  },
  data() {
    return {
      listLoading: false,
      tableKey: 0,
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
    handleFilter: function() {
      this.listQuery.pageIndex = 1;
      this.getList();
    },
    handleCreate: function() {
      this.dialogStatus = "create";
      this.dialogFormVisible = true;
      this.temp = {};
    },
    handleUpdate: function(e) {
      this.dialogStatus = "update";
      this.dialogFormVisible = true;
      this.temp = Object.assign({}, e);
    },
    
    handleDelete: function(e) {
      this.dialogPvVisible = true;
      this.temp = Object.assign({}, e);
    },
    getList() {
      this.listLoading = true;
      let channelName = this.listQuery.channelName;
      let pageIndex = this.listQuery.pageIndex;
      getChannelList(channelName, pageIndex)
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
      getChannelUpdate(this.temp).then(()=>{
        this.getList()
        this.dialogFormVisible = false
        this.showSuccess()
      })
    },
    createData() {
      getChannelAdd(this.temp).then(()=>{
        this.getList()
        this.dialogFormVisible = false
        this.showSuccess()
      })
    },
    deleteData() {
      getChannelDelete(this.temp).then(()=>{
        this.getList()
        this.dialogPvVisible = false
        this.showSuccess()
      })
    },
    showSuccess(){
      this.$message({
          message: '修改成功',
          type: 'success',
          duration:1000
        });
    }
  }
};
</script>

