package pl.edu.uj.projobj.zadanie3_kotlin.singletons

import org.springframework.stereotype.Service
import pl.edu.uj.projobj.zadanie3_kotlin.models.UserModel

@Service
object AuthServiceSingletonEager {
    private val registeredUsers: MutableMap<String, String>  = mutableMapOf("1" to "1")

    fun addUser(user: UserModel): Boolean {
        if (registeredUsers.containsKey(user.login)) return false
        registeredUsers[user.login] = user.password
        return true
    }

    fun authenticate(user: UserModel): Boolean {
        if (!registeredUsers.containsKey(user.login)) return false
        return registeredUsers[user.login].equals(user.password)
    }
}