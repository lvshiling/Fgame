<template>
    <div class="app-container">
        <div class="filter-container">
            <el-input placeholder="用户名" v-model="listQuery.userName" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
            <el-select v-model="listQuery.privilege" placeholder="角色" clearable style="width: 120px" class="filter-item">
                <el-option v-for="item in privilegeList" :key="item.key" :label="item.name" :value="item.key"/>
            </el-select>
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
            <el-table-column label="用户ID" align="center" width="65">
                <template slot-scope="scope">
                    <span>{{ scope.row.userId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="用户名" width="150px" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.userName}}</span>
                </template>
            </el-table-column>
            <el-table-column label="用户角色" width="150px">
                <template slot-scope="scope">
                <span>{{ scope.row.privilegeName}}</span>
                </template>
            </el-table-column>
            <el-table-column label="渠道" width="150px">
                <template slot-scope="scope">
                <span>{{ getChannelName(scope.row.channelId) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="平台" min-width="150px">
                <template slot-scope="scope">
                <span>{{ getPlatName(scope.row.platformId) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作" align="center" width="280" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                <el-button type="primary" size="mini" @click="handleUpdate(scope.row)">编辑</el-button>
                <el-button type="primary" size="mini" style="width:70px;" @click="handlePwd(scope.row)">重置密码</el-button>
                <el-button size="mini" type="danger" @click="handleDelete(scope.row)">删除</el-button>
                </template>
            </el-table-column>
            </el-table>

        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>

        <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
        <el-form ref="dataForm" :model="temp" label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">
            <el-form-item label="用户名">
                <el-input v-model="temp.userName"/>
            </el-form-item>
            <el-form-item v-if="dialogStatus=='create'" label="密码">
                <el-input v-model="temp.password" type="password"/>
            </el-form-item>
            <el-form-item label="角色">
                <el-select v-model="temp.privilegeid" class="filter-item" placeholder="选择角色">
                    <el-option v-for="item in privilegeList" :key="item.key" :label="item.name" :value="item.key"/>
                </el-select>
            </el-form-item>

            <el-form-item v-if="checkHasChannel(temp)" label="渠道">
                <el-select v-model="temp.channelId" class="filter-item" placeholder="选择渠道" @change="handleChannelChange">
                    <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
                </el-select>
            </el-form-item>

            <el-form-item v-if="checkHasPlatform(temp)" label="平台">
                <el-select v-model="temp.platformId" class="filter-item" placeholder="选择平台">
                    <el-option v-for="item in tempPlatFormList" :key="item.platformId" :label="item.platformName" :value="item.platformId"/>
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
              是否确认删除用户：{{temp.userName}}
          </div>
          <span slot="footer" class="dialog-footer">
            <el-button @click="dialogPvVisible = false">取消</el-button>
            <el-button type="danger" @click="deleteData">删除</el-button>
          </span>
        </el-dialog>

        <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogPassWorldVisible">
          <el-form ref="dataForm" :model="temp" label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">
              <el-form-item  label="密码">
                  <el-input v-model="temp.password" type="password"/>
              </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
              <el-button @click="dialogPassWorldVisible = false">取消</el-button>
              <el-button type="primary" @click="changepass">确定</el-button>
          </div>
        </el-dialog>
    </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import {
  getUserList,
  saveUserInfo,
  deleteUserInfo,
  changePassword,
  childPrivilege
} from "@/api/user";
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { PrivilegeType, PrivilegeTypeMap } from "@/types/privilege";
export default {
  name: "UserList",
  directives: {
    waves
  },
  created() {
    this.initMetaData();
    this.getList();
    this.privilegeInstance = PrivilegeType;
    childPrivilege().then(res => {
      this.privilegeList = res;
    });
    // this.privilegeList = PrivilegeTypeMap;
  },
  data() {
    return {
      listLoading: false,
      tableKey: 0,
      total: 0,
      listQuery: {
        pageIndex: 1,
        userName: "",
        privilege: ""
      },
      textMap: {
        update: "编辑",
        create: "添加",
        psd: "重置密码"
      },
      dialogStatus: "",
      dialogPvVisible: false,
      dialogFormVisible: false,
      dialogPassWorldVisible: false,
      temp: {},
      privilegeList: [],
      list: [],
      channelList: [],
      platformList: [],
      tempPlatFormList: []
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
        userName: "",
        password: "",
        privilegeid: undefined,
        channelId: undefined,
        platformId: undefined
      };
    },
    handleUpdate: function(e) {
      this.dialogStatus = "update";
      this.dialogFormVisible = true;
      this.temp = this.copyData(e);
    },
    handlePwd: function(e) {
      this.dialogStatus = "psd";
      this.dialogPassWorldVisible = true;
      this.temp = this.copyData(e);
    },
    handleDelete: function(e) {
      this.dialogPvVisible = true;
      this.temp = this.copyData(e);
    },
    handleChannelChange: function(e) {
      if (e) {
        this.temp.platformId = undefined;
        this.tempPlatFormList = this.findPlatFormList(e);
        if (this.tempPlatFormList && this.tempPlatFormList.length > 0) {
          this.temp.platformId = this.tempPlatFormList[0].platformId;
        }
      }
    },
    // handlePlatformChange: function(e) {
    //   // let platformInfo = this.findPlatFormItem(e);
    //   // if (platformInfo) {
    //   //   this.temp.channelId = platformInfo.channelId;
    //   // }
    // },
    getList() {
      this.listLoading = true;
      let username = this.listQuery.userName;
      let privilege = -1;
      if (this.listQuery.privilege != "") {
        privilege = parseInt(this.listQuery.privilege);
      }
      let pageIndex = this.listQuery.pageIndex;
      getUserList(username, privilege, pageIndex)
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
      let checkFlag = this.checkAddUpdate()
      if(!checkFlag){
        
        return
      }
      saveUserInfo(this.temp).then(() => {
        this.getList();
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    createData() {
      let checkFlag = this.checkAddUpdate()
      if(!checkFlag){
        return
      }
      saveUserInfo(this.temp).then(() => {
        this.getList();
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    deleteData() {
      deleteUserInfo(this.temp).then(() => {
        this.getList();
        this.dialogPvVisible = false;
        this.showSuccess();
      });
    },
    changepass() {
      changePassword(this.temp).then(() => {
        this.getList();
        this.dialogPassWorldVisible = false;
        this.showSuccess();
      });
    },
    initMetaData() {
      getAllChannel().then(res => {
        this.channelList = res.itemArray;
      });
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
        this.tempPlatFormList = this.platformList;
      });
    },
    getChannelName(channelId) {
      let channelName = this.channelList.find(n => {
        return n.channelId == channelId;
      });
      if (channelName && channelName.channelName) {
        return channelName.channelName;
      }
      return "";
    },
    getPlatName(platid) {
      let platformName = this.platformList.find(n => {
        return n.platformId == platid;
      });
      if (platformName && platformName.platformName) {
        return platformName.platformName;
      }
      return "";
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
    copyData(item) {
      let result = Object.assign({}, item);
      if (result.channelId == 0) {
        result.channelId = undefined;
      }
      if (result.platformId == 0) {
        result.platformId = undefined;
      }
      return result;
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
    },
    checkAddUpdate(){
      if(this.checkHasChannel(this.temp)){
        if(!this.temp.channelId){
          this.showError("请选择渠道")
          return false
        }
      }

      if(this.checkHasPlatform(this.temp)){
        if(!this.temp.platformId){
          this.showError("请选择平台")
          return false
        }
      }
      return true
    },
    checkHasChannel(temp){
      if(temp.privilegeid == this.privilegeInstance.PrivilegeLevelPlatform 
      || temp.privilegeid == this.privilegeInstance.PrivilegeLevelChannel 
      || temp.privilegeid == this.privilegeInstance.PrivilegeLevelKeFu 
      || temp.privilegeid == this.privilegeInstance.PrivilegeLevelCommonKeFu 
      || temp.privilegeid == this.privilegeInstance.PrivilegeLevelMinitor 
      || temp.privilegeid == this.privilegeInstance.PrivilegeLevelGaoJiKeFu
      || temp.privilegeid == this.privilegeInstance.PrivilegeLevelNeiGua
      || temp.privilegeid == this.privilegeInstance.PrivilegeLevelGs){
        return true
      }
      return false
    },
    checkHasPlatform(temp){
      if(temp.privilegeid == this.privilegeInstance.PrivilegeLevelPlatform 
      || temp.privilegeid == this.privilegeInstance.PrivilegeLevelKeFu 
      || temp.privilegeid == this.privilegeInstance.PrivilegeLevelCommonKeFu 
      || temp.privilegeid == this.privilegeInstance.PrivilegeLevelMinitor 
      || temp.privilegeid == this.privilegeInstance.PrivilegeLevelGaoJiKeFu
      || temp.privilegeid == this.privilegeInstance.PrivilegeLevelNeiGua
      || temp.privilegeid == this.privilegeInstance.PrivilegeLevelGs){
        return true
    }
    return false
    }
  }
};
</script>

