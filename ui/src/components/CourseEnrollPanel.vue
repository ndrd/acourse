<template lang="pug">
md-whiteframe(md-elevation='1')
  md-layout(md-column)
    md-layout(v-if='!course.purchase', md-align='center')
      span(v-if="price <= 0")
        h2 FREE
      span(v-else, md-align='center')
        h2 ฿ {{ price }}
    md-layout(v-if="isAuth", md-align='center')
      md-button.md-raised.md-primary(v-if='!course.purchase', @click.native='apply') Enroll
      md-button.md-raised(v-else, disabled) Wait for Approve
    md-layout(v-else, md-align='center')
      md-button.md-raised(disabled) Sign In to Enroll
  .ui.stackable.equal.width.grid
    .row(v-if="isAuth")
      .ui.blue.button(style="width: 200px;", :class="{loading: applying}", @click="apply", v-if="!course.purchase") Enroll
  EnrollModal(ref="enrollModal", :course="course")
</template>

<script>
import { Course, Document, Auth } from 'services'
import EnrollModal from './EnrollModal'

export default {
  components: {
    EnrollModal
  },
  props: {
    course: {
      type: Object,
      required: true
    }
  },
  data () {
    return {
      applying: false
    }
  },
  subscriptions () {
    return {
      isAuth: Auth.currentUser().map((x) => !!x)
    }
  },
  methods: {
    apply () {
      if (this.applying) return

      if (!this.price) {
        this.applying = true
        Course.enroll(this.course.id, {})
          .finally(() => { this.applying = false })
          .subscribe(
            () => {
              Document.openSuccessModal('Success', 'You have enrolled to this course.')
              Course.fetch(this.course.id)
            }
          )
      } else {
        this.$refs.enrollModal.show()
      }
    }
  },
  computed: {
    price () {
      if (this.course.discount) return this.course.discountedPrice
      return this.course.price || 0
    }
  }
}
</script>
