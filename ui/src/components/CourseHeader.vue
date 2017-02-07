<template lang="pug">
md-layout(md-column)
  md-layout(md-align='center')
    img(:src='course.photo', style='max-height: 300px;')
  md-layout(md-align='center')
    h1 {{ course.title }}
  md-layout(md-align='center', style='margin-bottom: 1rem;')
    Avatar(:src='course.owner.photo', size='mini', style='margin: initial;')
    h3(style='padding-left: 1rem;') {{ course.owner.name || 'Anonymous' }}
</template>

<script>
import Avatar from './Avatar'

export default {
  components: {
    Avatar
  },
  props: {
    course: {
      type: Object,
      required: true
    }
  }
}
</script>
