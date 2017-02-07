<template lang="pug">
div(:class='{loading: !course}')
  div(style='padding-bottom: 1.5rem;')
    router-link.md-headline(to='/') Courses
    router-link.md-title(:to='`/course/${courseId}`', :tag="$route.name === 'courseView' && 'div' || 'a'", active-class='active', exact='') {{ course && course.title || courseId }}
    .active.section(v-show="$route.name === 'courseEdit'") Edit
    .active.section(v-show="$route.name === 'courseNew'") New
    .active.section(v-show="$route.name === 'courseAssignments'") Assignments
    .active.section(v-show="$route.name === 'courseAttend'") Attendants
    .active.section(v-show="$route.name === 'courseAssignmentEdit'") Edit Assignment
  router-view
</template>

<script>
import { Course, Document } from 'services'

export default {
  data () {
    return {
      courseId: this.$route.params.id
    }
  },
  subscriptions () {
    return {
      course: Course.get(this.courseId)
        .do((course) => Document.setCourse(course))
    }
  },
  created () {
    Course.fetch(this.courseId)
      .subscribe(null, () => {
        this.$router.replace('/')
      })
  }
}
</script>
