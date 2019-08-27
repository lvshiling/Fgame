<template>
  <div class="dashboard-container">
    <component :is="currentRole"/>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import adminDashboard from './admin'
import editorDashboard from './editor'
import serviceDashboard from './service'

export default {
  name: 'Dashboard',
  components: { adminDashboard, editorDashboard,serviceDashboard },
  data() {
    return {
      currentRole: 'adminDashboard',
    }
  },
  computed: {
    ...mapGetters([
      'roles'
    ])
  },
  created() {
    if (this.roles.includes('minitor')) {
      this.currentRole = 'editorDashboard'
    }
    if (this.roles.includes('super_channel_service')) {
      this.currentRole = 'serviceDashboard'
    }
  }
}
</script>
