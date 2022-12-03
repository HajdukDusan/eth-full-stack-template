<template>
  <div id="app">
    <Navbar/>

    <div class="d-flex content">
        <div class="col-2 px-0 bg-light">
          <Sidebar/>
        </div>
        <div class="col-10 px-3 pt-3">
          <router-view></router-view>
        </div>
    </div>
  </div>
</template>

<script>
import Navbar from './components/Navbar.vue';
import Sidebar from './components/Sidebar.vue';

export default {
  name: 'App',
  components: {
    Navbar,
    Sidebar
  }
}
</script>

<style>
.content {
  height: 100vh;
}
</style>