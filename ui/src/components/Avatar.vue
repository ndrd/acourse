<template lang="pug">
md-avatar(:class='`a-${size}`')
  img(:src='src || `/static/icons/ic_face_black_48px.svg`')
</template>

<style lang="scss" scoped>
  .md-avatar {
    &.a-tiny {
      width: 32px;
      height: 32px;
      min-width: 32px;
      min-height: 32px;
    }

    &.a-mini {
      width: 64px;
      height: 64px;
      min-width: 64px;
      min-height: 64px;
    }

    &.a-small {
      width: 120px;
      height: 120px;
      min-width: 120px;
      min-height: 120px;
    }

    &.a-medium {
      width: 200px;
      height: 200px;
      min-width: 200px;
      min-height: 200px;
    }
  }

  img {
    object-fit: cover;
    object-position: top;
  }
</style>

<script>
export default {
  props: {
    src: {},
    size: {
      type: String,
      required: true
    }
  }
}
</script>
