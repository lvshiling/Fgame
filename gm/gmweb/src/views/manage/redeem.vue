<template>
    <div class="app-container">
        <div>
            <div class="filter-container">
                <el-select v-model="listQuery.platformId" placeholder="平台" clearable style="width: 160px" class="filter-item" @change="handlePlatformChange">
                    <el-option v-for="item in platformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
                </el-select>
                <el-input placeholder="礼包名" v-model="listQuery.name" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>

                <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
                <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">添加</el-button>
            </div>
        </div>

        <el-table
            :key="chatSetKey"
            :data="chatSetList"
            v-loading="listLoading"
            border
            fit
            highlight-current-row
            style="width: 100%;margin-top:15px;">
            <el-table-column label="配置Id" align="center" width="100px">
                <template slot-scope="scope">
                  <router-link :to="'/gameplayer/redeemcode/'+scope.row.id" class="link-type">
                    <span>{{ scope.row.id }}</span>
                  </router-link>
                    <!-- <span>{{ scope.row.id }}</span> -->
                </template>
            </el-table-column>
           <el-table-column label="礼包名称" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.giftBagName }}</span>
                </template>
            </el-table-column>
            
           <el-table-column label="兑换码数量" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.redeemNum }}</span>
                </template>
            </el-table-column>
            <el-table-column label="兑换码使用次数" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.redeemUseNum }}</span>
                </template>
            </el-table-column>
            <el-table-column label="个人使用次数" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.redeemPlayerUseNum }}</span>
                </template>
            </el-table-column>
            <el-table-column label="全服使用次数" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.redeemServerUseNum }}</span>
                </template>
            </el-table-column>
            <el-table-column label="sdk类别" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ getSdkTypeName(scope.row.sdkTypes) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="发送方式" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.sendType | parseSendType }}</span>
                </template>
            </el-table-column>
            <el-table-column label="开始时间" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.startTime | parseTime}}</span>
                </template>
            </el-table-column>
            <el-table-column label="结束时间" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.endTime | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column label="最低等级" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.minPlayerLevel }}</span>
                </template>
            </el-table-column>
            <el-table-column label="最低VIP等级" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.minVipLevel }}</span>
                </template>
            </el-table-column>
            <el-table-column label="是否生成" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.createFlag | parseYesOrNo}}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作" align="center" width="260" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                <el-button type="primary" size="mini" @click="handleEdit(scope.row)">查看</el-button>
                <el-button size="mini" type="danger" @click="handleDelete(scope.row)">删除</el-button>
                <el-button v-if="scope.row.createFlag == 1" type="primary" size="mini" @click="handledown(scope.row)">下载</el-button>
                </template>
            </el-table-column>
        </el-table>
        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>

        <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
            <el-form ref="dataForm" :model="temp" label-position="left" label-width="100px" style="width: 400px; margin-left:50px;">
                <el-form-item label="平台">
                    <el-select v-model="temp.sdkTypes" multiple placeholder="平台" :disabled="dialogStatus=='update'" style="width: 160px" class="filter-item"  @change="handleItemPlatformChange">
                        <el-option v-for="item in platformList" :key="item.sdkType" :label="item.platformName" :value="item.sdkType" />
                    </el-select>
                </el-form-item>

                <el-form-item label="礼包名称">
                    <el-input v-model="temp.giftBagName" :disabled="dialogStatus=='update'"/>
                </el-form-item>
                <el-form-item label="礼包说明">
                    <el-input v-model="temp.giftBagDesc" :disabled="dialogStatus=='update'"/>
                </el-form-item>
                <el-form-item label="礼包内容">
                    <el-input v-model="temp.giftBagContent" :disabled="dialogStatus=='update'"/>
                </el-form-item>
                <el-form-item label="兑换码数量">
                    <el-input type="number" v-model="temp.redeemNum" :disabled="dialogStatus=='update'"/>
                </el-form-item>
                <el-form-item label="兑换码使用次数">
                    <el-input type="number" v-model="temp.redeemUseNum" :disabled="dialogStatus=='update'"/>
                </el-form-item>
                <el-form-item label="个人使用次数">
                    <el-input type="number" v-model="temp.redeemPlayerUseNum" :disabled="dialogStatus=='update'"/>
                </el-form-item>
                <el-form-item label="全服使用次数">
                    <el-input type="number" v-model="temp.redeemServerUseNum" :disabled="dialogStatus=='update'"/>
                </el-form-item>
                <el-form-item label="发放方式">
                    <el-select v-model="temp.sendType" placeholder="发放方式" :disabled="dialogStatus=='update'">
                        <el-option v-for="item in sendTypeList" :key="item.key" :label="item.name" :value="item.key" />
                    </el-select>
                </el-form-item>
                <el-form-item label="生效范围" :disabled="dialogStatus=='update'">
                    <el-date-picker v-model="temp.startEnd" type="datetimerange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期" :disabled="dialogStatus=='update'">
                </el-date-picker>
                </el-form-item>
                <el-form-item label="最低等级" >
                    <el-input type="number" v-model="temp.minPlayerLevel" :disabled="dialogStatus=='update'"/>
                </el-form-item>
                <el-form-item label="最低VIP等级" >
                    <el-input type="number" v-model="temp.minVipLevel" :disabled="dialogStatus=='update'"/>
                </el-form-item>
            </el-form>

            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogFormVisible = false">取消</el-button>
                <el-button v-if="dialogStatus=='create'" type="primary" @click="createData">创建</el-button>
                <el-button v-else-if="temp.createFlag == 0" type="primary" @click="updateData">生成</el-button>
            </div>
        </el-dialog>

        <el-dialog :visible.sync="dialogPvVisible" title="是否确认删除" width="30%">
        <div>
                是否确认删除：
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
import { parseTime } from "@/utils/index";
import { getCenterPlatList } from "@/api/center";
import { getCenterServerList } from "@/api/center";
import { getAllPlatformList } from "@/api/platform";
import { sendType } from "@/types/redeem";
import { yesOrNoList } from "@/types/public";
import { getAllSdkType } from "@/api/center";
import {
  addRedeem,
  getRedeemList,
  codeRedeem,
  deleteRedeem,
  getRedeemCodeListExport
} from "@/api/redeem";
import { sdkTypeList } from "@/types/center";
export default {
  name: "Redeem",
  directives: {
    waves
  },
  filters: {
    parseTime: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
    },
    parseSendType: function(value) {
      let item = sendType[value - 1];
      if (item) {
        return item.name;
      }
      return "";
    },
    parseYesOrNo: function(value) {
      if (value == 1) {
        return "是";
      }
      return "否";
    },
    parseSdkType: function(value) {
      if (value.length == 0) {
        return "全渠道";
      }
      let msg = "";
      for (let i = 0, len = value.length; i < len; i++) {
        let item = sdkTypeList[value[i] - 1];
        if (item) {
          if (!msg) {
            msg = item.name;
          } else {
            msg = msg + "," + item.name;
          }
        }
      }
      return msg;
    }
  },
  created() {
    this.initMetaData();
    getAllSdkType().then(res =>{
        this.skdTypeArray = res.itemArray
    });
  },
  mounted() {
    //   this.websocket.close()
  },
  beforeDestroy() {},
  destroyed() {},
  data() {
    return {
      chatSetKey: 1,
      listLoading: false,
      total: 0,
      chatSetList: [],
      dialogFormVisible: false,
      dialogPvVisible: false,
      dialogStatus: undefined,
      textMap: {
        update: "编辑",
        create: "添加"
      },
      listQuery: {
        centerPlatformId: undefined,
        centerServerId: undefined,
        centerServerId: undefined,
        pageIndex: 0
      },
      platformList: [],
      serverList: [],
      temp: {
        startTime: undefined,
        endTime: undefined,
        centerPlatformId: undefined,
        centerServerId: undefined,
        platformArray: []
      },
      tempServerList: [],
      sendTypeList: [],
      skdTypeArray:[]
    };
  },
  methods: {
    initMetaData() {
      this.sendTypeList = sendType;
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
      });
      // this.loadData();
    },
    handleFilter(e) {
      this.listQuery.pageIndex = 1;
      this.loadData();
    },
    handleCreate(e) {
      this.temp = {
        startTime: undefined,
        endTime: undefined,
        centerPlatformId: undefined,
        centerServerId: undefined,
        platformArray: [],
        sdkTypes: []
      };
      this.dialogFormVisible = true;
      this.dialogStatus = "create";
      this.tempServerList = [];
    },
    handleEdit(e) {
      this.dialogFormVisible = true;
      this.dialogStatus = "update";
      this.temp = Object.assign({}, e);
      this.temp.startEnd = [
        new Date(this.temp.startTime),
        new Date(this.temp.endTime)
      ];
      //   this.temp.startTime = new Date("2000-01-01 " + this.temp.startTime);
      //   this.temp.endTime = new Date("2000-01-01 " + this.temp.endTime);
      //   let item = this.findPlatFormItemByCenterId(this.temp.centerPlatformId);
      //   if (item) {
      //     this.temp.platformId = item.platformId;
      //   }
      //   getCenterServerList(this.temp.centerPlatformId).then(res => {
      //     this.tempServerList = res.itemArray;
      //   });
    },
    handleDelete(e) {
      this.dialogPvVisible = true;
      this.temp = Object.assign({}, e);
    },
    handleCurrentChange(e) {
      this.listQuery.pageIndex = e;
      this.loadData();
    },
    handledown(e){
      this.temp = Object.assign({}, e);
      getRedeemCodeListExport(this.temp)
        .then(res => {
          this.listLoading = false;
          if (!res) {
                return
            }
            let name = this.temp.giftBagName + "_兑换码.xlsx"
            let url = window.URL.createObjectURL(new Blob([res.data]))
            let link = document.createElement('a')
            link.style.display = 'none'
            link.href = url
            link.setAttribute('download', name)

            document.body.appendChild(link)
            link.click()
        })
        .catch(() => {
          this.listLoading = false;
        });
    },
    handlePlatformChange(e) {
      if (!e || e == "" || e == undefined) {
        this.listQuery.centerPlatformId = undefined;
        this.listQuery.centerServerId = undefined;
        return;
      }
      this.listQuery.centerServerId = undefined;
      let platformItem = this.findPlatFormItem(this.listQuery.platformId);

      if (platformItem) {
        this.listQuery.sdkType = platformItem.sdkType;
        //     this.listQuery.centerPlatformId = platformItem.centerPlatformId;
        //     getCenterServerList(platformItem.centerPlatformId).then(res => {
        //       this.serverList = res.itemArray;
        //     });
      }
    },
    handleItemPlatformChange(e) {
      this.temp.centerServerId = undefined;
      let platformItem = this.findPlatFormItem(this.temp.platformId);
      if (platformItem) {
        this.temp.centerPlatformId = platformItem.centerPlatformId;
        this.temp.sdkType = platformItem.sdkType;
        getCenterServerList(this.temp.centerPlatformId).then(res => {
          this.tempServerList = res.itemArray;
        });
      }
    },
    findPlatFormName(platformId) {
      let platform = this.platformList.find(n => {
        return n.centerPlatformId == platformId;
      });
      if (platform) {
        return platform.platformName;
      }
      return undefined;
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
    findPlatFormItemByCenterId(centerPlatformId) {
      let platform = this.platformList.find(n => {
        return n.centerPlatformId == centerPlatformId;
      });
      if (platform) {
        return platform;
      }
      return undefined;
    },
    createData(e) {
      if (!this.temp.startEnd || this.temp.startEnd.length != 2) {
        Message({
          message: "请选生效时间",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      if (this.temp.startEnd && this.temp.startEnd.length == 2) {
        this.temp.startTime = this.temp.startEnd[0].valueOf();
        this.temp.endTime = this.temp.startEnd[1].valueOf();
      }
      addRedeem(this.temp).then(res => {
        this.dialogFormVisible = false;
        this.showSuccess();
        this.loadData();
      });
    },
    updateData(e) {
      if (this.temp.startEnd && this.temp.startEnd.length == 2) {
        this.temp.startTime = this.temp.startEnd[0].valueOf();
        this.temp.endTime = this.temp.startEnd[1].valueOf();
      }
      codeRedeem(this.temp).then(res => {
        this.dialogFormVisible = false;
        this.showSuccess();
        this.loadData();
      });
    },
    deleteData(e) {
      deleteRedeem(this.temp).then(res => {
        this.dialogPvVisible = false;
        this.showSuccess();
        this.loadData();
      });
    },
    loadData() {
      if(!this.listQuery.sdkType){
        this.showFailMsg("请选择平台")
        return
      }
      this.listLoading = true;
      getRedeemList(this.listQuery).then(res => {
        this.chatSetList = res.itemArray;
        this.total = res.total;
        this.listLoading = false;
      });
    },
    showSuccess() {
      this.$message({
        message: "修改成功",
        type: "success",
        duration: 1000
      });
    },  
    showFailMsg(msg) {
      this.$message({
        message: msg,
        type: "error",
        duration: 1000
      });
    },
    getSdkTypeName(value){
        for(let i=0,len=this.skdTypeArray.length;i<len;i++){
            if(this.skdTypeArray[i].key == value){
                return this.skdTypeArray[i].name
            }
        }
        return value
    }
  }
};
</script>

