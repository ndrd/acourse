<template lang="pug">
  .ui.small.modal
    .header วิธีการลงทะเบียน {{ course.title }}
    .content
      div(v-html="detail")
      br
      .ui.form
        .field
          label ใส่จำนวนเงินที่โอน (บาท)
          input(type='number', v-model.number='price')
      br
      .ui.fluid.green.button(@click='enroll') Upload and Enroll
</template>

<script>
import { Document, Course } from 'services'

export default {
  props: {
    course: {
      type: Object,
      required: true
    }
  },
  data () {
    return {
      price: 0,
      url: '',
      code: ''
    }
  },
  computed: {
    calcPrice () {
      const price = this.course.discount ? this.course.discountedPrice : this.course.price || 0
      return price <= 0 ? 0 : price
    },
    detail () {
      if (!this.course.enrollDetail) return ''
      return this.marked(this.course.enrollDetail)
    }
  },
  methods: {
    show () {
      this.price = this.calcPrice
      $(this.$el).modal('show')
    },
    enroll () {
      Document.uploadModal.open('image/*')
        .do((file) => { this.url = file.downloadURL })
        .flatMap(() => Course.enroll(this.course.id, { url: this.url, price: this.price, code: this.code }))
        .subscribe(
          () => {
            Course.fetch(this.course.id)
            Document.openSuccessModal('Success', 'Your enroll request success!.')
          },
          (err) => {
            Document.openErrorModal('Error', err && err.message || err)
          }
        )
    }
  }
}
</script>
