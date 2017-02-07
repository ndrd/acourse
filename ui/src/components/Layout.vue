<template lang="pug">
  div
    Navbar
    md-layout#view
      md-layout(md-hide-xsmall, md-flex-small='5', md-flex-medium='10', md-flex-large='15', md-flex-xlarge='20')
      md-layout
        router-view
      md-layout(md-hide-xsmall, md-flex-small='5', md-flex-medium='10', md-flex-large='15', md-flex-xlarge='20')
</template>

<style scoped>
  #view {
    padding-top: 1rem;
    padding-bottom: 1rem;
  }
</style>

<script>
import Navbar from './Navbar'

export default {
  components: {
    Navbar
  }
}
</script>
