<template>
   <a-layout class="container">
      <Nav></Nav>
      <a-layout>
        <a-layout-header class="headerBtn">
            <Header></Header>
        </a-layout-header>
        <a-layout-content>
            <router-view>

            </router-view>
        </a-layout-content>
        <a-layout-footer>
            <Footer></Footer>
        </a-layout-footer>
      </a-layout>
    </a-layout>
</template>

<script>

import Nav from '../components/mgt/Nav'
import Footer from '../components/mgt/Footer'
import Header from '../components/mgt/Header'

export default {
    components: {Nav, Header, Footer}
}

</script>

<style scoped>
.container {
    height: 100%;
}
.headerBtn {
    display: flex;
    justify-content: flex-end;
    align-items: center;
}
</style>
