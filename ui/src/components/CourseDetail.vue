<template lang="pug">
  .ui.segment
    h4.ui.header คำอธิบาย
    span เริ่มเรียน: {{ course.start | date('DD/MM/YYYY') }}
    .ui.divider
    p.description(v-html='course.description')
</template>

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
