<template lang="pug">
div
  md-toolbar.md-dense
    router-link(to='/')
      md-icon(:md-src=`require('assets/acourse.svg')`, style='margin-right: 1rem;')
    h2.md-title(style='flex: 1;') Acourse
    md-menu(v-if='currentUser', md-direction='bottom right', ref='profileMenu', md-size='5')
      div(md-menu-trigger)
      md-menu-content
        div(style='margin: .5rem 1.5rem;')
          md-layout
            Avatar(:src='currentUser.photo', size='mini', style='margin: initial;')
            md-layout(md-column, style='margin: .5rem 0 .5rem 1rem;')
              md-layout
                | {{ currentUser.username }}
              md-layout
                a(href='#', @click.prevent='signOut') SIGN OUT
    Avatar(v-if='currentUser', :src='currentUser.photo', size='tiny', @click.native='$refs.profileMenu.open')
    router-link(v-else, tag='md-button', to='/signin') SIGN IN
  .ui.borderless.top.fixed.menu
    .right.menu
      .ui.dropdown.item(ref='dropdownAdmin', v-if='currentUser && currentUser.role && currentUser.role.admin')
        | Admin
        i.dropdown.icon
        .menu
          router-link.item(to='/admin/course') Course
          router-link.item(to='/admin/payment') Payment
          router-link.item(to='/admin/payment/history') Payment History
      .ui.dropdown.item(ref='dropdownUser', v-if='currentUser', style='padding-top: 0.5rem; padding-bottom: 0.5rem;')
        UserAvatar(:user='currentUser')
        i.dropdown.icon
        .menu
          router-link.item(to='/profile') Profile
          a.item(@click='signOut') Sign Out
</template>

<script>
import { Auth, Me } from 'services'
import UserAvatar from './UserAvatar'
import Avatar from './Avatar'

export default {
  components: {
    UserAvatar,
    Avatar
  },
  subscriptions () {
    return {
      currentUser: Me.get()
    }
  },
  methods: {
    signOut () {
      Auth.signOut().subscribe(() => {
        this.$nextTick(() => {
          this.$router.push('/')
        })
      })
    }
  }
}
</script>
