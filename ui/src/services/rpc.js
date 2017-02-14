import Firebase from './firebase'
import { Observable } from 'rxjs/Observable'

const API_URL = process.env.API_URL

const ifThen = (x, y) => !x || y

const getConfig = () => Firebase.getToken()
  .first()
  .map((token) => {
    const x = {
      headers: {},
      isAuth () {
        return !!token
      }
    }
    if (token) {
      x.headers.Authorization = `Bearer ${token}`
    }
    return x
  })

const getResponse = (res) => res && res.data

const justNull = Observable.of(null)

const invoke = (url, data, requireAuth) => getConfig()
  .flatMap((config) =>
    ifThen(requireAuth, config.isAuth())
      ? Observable.fromPromise(axios.post(API_URL + url, data || {}, config))
      : justNull)
  .map(getResponse)

export default {
  invoke
}
