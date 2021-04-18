<template>
  <a-button type="danger"
            @click="logout">退出</a-button>
</template>

<script>

export default {
  methods: {
    logout () {
      window.sessionStorage.clear('token')
      this.$router.push('login')
    }
  }
}

</script>

<style scoped>
.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>