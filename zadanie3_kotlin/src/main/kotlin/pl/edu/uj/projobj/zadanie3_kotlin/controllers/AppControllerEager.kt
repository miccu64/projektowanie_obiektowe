package pl.edu.uj.projobj.zadanie3_kotlin.controllers

import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.server.ResponseStatusException
import pl.edu.uj.projobj.zadanie3_kotlin.models.DataModel
import pl.edu.uj.projobj.zadanie3_kotlin.models.UserModel
import pl.edu.uj.projobj.zadanie3_kotlin.singletons.AuthServiceSingletonEager

@RestController
@RequestMapping("eager")
class AppControllerEager {
    @Autowired
    private lateinit var authService: AuthServiceSingletonEager
    private val data: DataModel = DataModel(mapOf(1 to "val1", 2 to "val2", 3 to "val3"), listOf("someVal1", "someVal2"))

    @PostMapping("register")
    fun register(@RequestBody user: UserModel) {
        if (!authService.addUser(user)) throw ResponseStatusException(
            HttpStatus.CONFLICT, "User already exists"
        )
    }

    @GetMapping("data")
    fun getData(@RequestParam login: String, @RequestParam password: String): DataModel {
        if (!authService.authenticate(UserModel(login, password))) throw ResponseStatusException(
            HttpStatus.UNAUTHORIZED, "Wrong credentials"
        )
        return data
    }
}

