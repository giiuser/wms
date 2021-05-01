<template>
<main>
  <Menu />
  <div class="container">
      <router-view></router-view>
  </div>
  <Footer />
</main>
</template>

<script>
import Footer from './components/Footer';
import Menu from './components/Menu';

export default {
  name: 'App',

  components: {
    Footer,
    Menu,
  },

  data: () => ({
    //
  }),
};
</script>
