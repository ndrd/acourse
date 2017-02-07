<template lang="pug">
router-link(:to='`/course/${url}`', tag='div')
  md-card(md-with-hover)
    md-card-area
      md-card-media
        img(:src='course.photo')
        md-ink-ripple
      md-card-header
        .md-title {{ course.title | trim(45) }}
        .md-subhead
          span(v-if="course.type === 'video'") Video
          span.date(v-if="course.type === 'live'") Live start at {{ course.start | date('DD/MM/YYYY') }}
          span(v-if="course.type === 'ebook'") eBook
      md-card-content
        span {{ course.shortDescription }}
    md-card-content
      md-layout
        md-layout
          span.price(v-if='!hidePrice && !course.discount && (course.price <= 0 || !course.price)') FREE
          span.price(v-if='!hidePrice && !course.discount && course.price > 0', :class='{line: course.discount}') ฿ {{ course.price | money }}
          span.discount.price(v-if='!hidePrice && course.discount') &nbsp;฿ {{ course.discountedPrice | money }}
        md-layout(md-align='end')
          span
            md-icon person
            | &nbsp;{{ course.student }}
</template>

<style scoped>
  .md-card-media img {
    object-fit: cover;
    object-position: center center;
    height: 180px !important;
  }

  .md-card-area {
    height: 360px;
  }

  .md-card > .md-card-content {
    padding-top: 12px;
    padding-bottom: 12px;
    color: rgba(0, 0, 0, .50);
  }

  .price {
    font-size: 1.1em;
  }

  .discount.price {
    color: red;
  }

  .price.line {
    text-decoration: line-through;
    font-size: .9em !important;
  }
</style>

<script>
export default {
  props: {
    course: {
      type: Object,
      required: true
    },
    hidePrice: {
      type: Boolean,
      required: false
    }
  },
  computed: {
    url () {
      if (!this.course) return ''
      return this.course.url || this.course.id
    }
  }
}
</script>
