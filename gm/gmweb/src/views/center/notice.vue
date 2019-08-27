<template>
    <div class="app-container">
        <el-tabs type="border-card">
            <el-tab-pane label="默认公告">
                <CenterNoticeSet></CenterNoticeSet>
            </el-tab-pane>
            <el-tab-pane label="平台公告列表">
                <CenterNoticeList></CenterNoticeList>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import CenterNoticeList from './notice/noticelist.vue'
import CenterNoticeSet from './notice/defaultnotice.vue'
export default {
  name: "CenterNotice",
  components: {
    CenterNoticeList,
    CenterNoticeSet
  },
  directives: {
    waves
  },
  created() {},
  data() {
    return {};
  },
  methods: {}
};
</script>

