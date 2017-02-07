<template lang="pug">
md-whiteframe(md-elevation='1')
  .ui.embed(ref='video', data-source='youtube', :data-id='src')
</template>

<style scoped>
  .md-whiteframe {
    margin-top: 1rem;
    padding: 1rem;
  }
</style>


<script>
export default {
  props: ['src'],
  mounted () {
    this.$nextTick(() => {
      $(this.$refs.video).embed()
    })
  }
}
</script>
