<template>
   <!-- <Menu />
   <router-view></router-view> -->
   <!-- <fragment>
  <menu></menu> 
   </fragment> -->
<main>
  <div class="container">
      <router-view></router-view>
  </div>
</main>
<!-- <fragment>
<footer></footer>
</fragment> -->
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
