<template lang="pug">
  .ui.basic.segment(:class='{loading}' style='padding: 0; margin: 0; min-height: 100%;')
    router-view
    SuccessModal(ref='successModal')
    ErrorModal(ref='errorModal')
    UploadModal(ref='uploadModal')
    Notification
  </div>
</template>

<style>
  .hidden {
    display: none;
  }

  .text-center {
    text-align: center;
  }

  p.description {
    text-align: left;
    white-space: pre-wrap;
    word-break: break-word;
  }
</style>

<script>
  import { Loader, Document } from 'services'
  import SuccessModal from 'components/SuccessModal'
  import ErrorModal from 'components/ErrorModal'
  import UploadModal from 'components/UploadModal'
  import Notification from 'components/Notification'

  export default {
    components: {
      SuccessModal,
      ErrorModal,
      UploadModal,
      Notification
    },
    data () {
      return {
        loader: Loader.state
      }
    },
    computed: {
      loading () {
        return !!this.loader.value
      }
    },
    created () {
      Document.$successModal
        .subscribe(
          ({title, description}) => {
            this.$refs.successModal.show(title, description)
          }
        )
      Document.$errorModal
        .subscribe(
          ({title, description}) => {
            this.$refs.errorModal.show(title, description)
          }
        )
    },
    mounted () {
      Document.uploadModal = this.$refs.uploadModal
    }
  }
</script>
