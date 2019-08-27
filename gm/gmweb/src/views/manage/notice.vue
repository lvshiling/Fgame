<template>
    <div class="app-container">
        <el-tabs type="border-card">
            <el-tab-pane label="公告发送">
                <NoticeSet></NoticeSet>
            </el-tab-pane>
            <el-tab-pane label="发送日志">
              <NoticeList></NoticeList>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import NoticeSet from './notice/noticeset.vue'
import NoticeList from './notice/noticelist.vue'
export default {
  name: "Notice",
  components: {
    NoticeSet,
    NoticeList
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

