import Firebase from './firebase'
import User from './user'
import Auth from './auth'
import { Observable } from 'rxjs'

export default {
  list () {
    return Firebase.onArrayValue('course')
  },
  get (id) {
    return Firebase.onValue(`course/${id}`)
      .map((course) => ({id, ...course}))
  },
  create (data) {
    data.timestamp = Firebase.timestamp
    return Firebase.push('course', data)
      .map((snapshot) => snapshot.key)
      // .flatMap((key) =>
      //   User.me()
      //     .first()
      //     .flatMap((user) =>
      //       User.updateMe({
      //         ...user,
      //         course: {
      //           ...user.course,
      //           [key]: true
      //         }
      //       })
      //     )
      //     .map(() => key)
      // )
  },
  save (id, data) {
    return Firebase.update(`course/${id}`, data)
  },
  favorite (id) {
    return Auth.currentUser
      .first()
      .flatMap((user) => Firebase.set(`course/${id}/favorite/${user.uid}`, true))
  },
  unfavorite (id) {
    return Auth.currentUser
      .first()
      .flatMap((user) => Firebase.remove(`course/${id}/favorite/${user.uid}`))
  },
  join (id) {
    return Auth.currentUser
      .first()
      .flatMap((user) =>
        Observable.forkJoin(
          Firebase.set(`course/${id}/student/${user.uid}`, true),
          User.addCourseMe(id)
        )
      )
  },
  ownBy (userId) {
    const ref = Firebase.ref('course').orderByChild('owner').equalTo(userId)
    return Firebase.onArrayValue(ref)
  },
  sendMessage (id, text) {
    return Auth.currentUser
      .first()
      .flatMap((auth) => Firebase.push(`chat/${id}`, {
        u: auth.uid,
        m: text,
        t: Firebase.timestamp
      }))
  },
  messages (id) {
    const ref = Firebase.ref(`chat/${id}`).orderByKey()
    return Firebase.onChildAdded(ref)
  }
}
