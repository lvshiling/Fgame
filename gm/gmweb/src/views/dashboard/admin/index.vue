<template>
  <div class="dashboard-container">
    <div class="dashboard-text">你好:{{ name }}</div>
    <div class="dashboard-text">在线人数:{{ onLineStatic.onLineNum }}人</div>
    <div class="dashboard-text">累计充值:{{ orderStatic.totalAmount }}元</div>
    <div class="dashboard-text">今日充值:{{ orderStatic.todayAmount }}元</div>
    <!-- <div class="dashboard-text">roles:<span v-for="role in roles" :key="role">{{ role }}</span></div> -->
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { getCenterOrderStaticTotal } from "@/api/centerorder";
import { getLastOnLineReport } from "@/api/report";

export default {
  name: "Dashboard",
  data() {
    return {
      currentRole: "adminDashboard",
      orderStatic: {
        totalAmount:0,
        todayAmount:0
      },
      onLineStatic:{
        onLineNum:0
      }
    };
  },
  computed: {
    ...mapGetters(["name", "roles"])
  },
  created() {
    getCenterOrderStaticTotal().then(res => {
      this.orderStatic = res;
    });
    getLastOnLineReport().then(res =>{
      this.onLineStatic=res;
    })
  }
};
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>