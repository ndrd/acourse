<template lang="pug">
md-whiteframe(md-elevation='1')
  h4 Course Detail
  span Start: {{ course.start | date('DD/MM/YYYY') }}
  p.description(v-html='course.description')
</template>

<style scoped>
  .md-whiteframe {
    margin-top: 1rem;
    padding: 1rem 2rem;
  }
</style>

<script>
export default {
  props: {
    course: {
      type: Object,
      required: true
    }
  }
}
</script>
