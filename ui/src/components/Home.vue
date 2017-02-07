<template lang="pug">
div(:class='{loading: courses === false}')
  div(v-if='courses && courses.length')
    h1.text-center All Courses
    md-layout
      md-layout(md-flex='33', md-flex-xsmall='100', md-flex-small='50', v-for='x in courses')
        CourseCard(:course='x', style='margin: .5rem;')
  div(v-if='courses && !courses.length')
    .ui.message No course available yet!
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
