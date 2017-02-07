<template lang="pug">
div(:class='{loading: courses === false}')
  div(v-if='courses && courses.length')
    span.md-headline Courses
    md-layout
      md-layout(md-flex='33', md-flex-xsmall='100', md-flex-small='50', v-for='x in courses')
        CourseCard(:course='x', style='margin: .5rem;')
  div(v-if='courses && !courses.length')
    span.md-display-1 No course available yet!
</template>

<style scoped>

</style>

<script>
import { Course } from 'services'
import CourseCard from './CourseCard'

export default {
  components: {
    CourseCard
  },
  subscriptions () {
    return {
      courses: Course.list()
    }
  },
  created () {
    this.fetchCourses()
  },
  methods: {
    fetchCourses () {
      Course.fetchList()
    }
  }
}
</script>
