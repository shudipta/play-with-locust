# from locust import HttpLocust, TaskSet, task, between

# class UserBehaviour(TaskSet):
#     def on_start(self):
#         """ on_start is called when a Locust start before any task is scheduled """
#         self.login()

#     def on_stop(self):
#         """ on_stop is called when the TaskSet is stopping """
#         self.logout()

#     def login(self):
#         self.client.post("/login", {"username":"ellen_key", "password":"education"})

#     def logout(self):
#         self.client.post("/logout", {"username":"ellen_key", "password":"education"})

#     @task(2)
#     def index(self):
#         self.client.get("/")

#     @task(1)
#     def profile(self):
#         self.client.get("/profile")

# class WebsiteUser(HttpLocust):
#     task_set = UserBehaviour
#     wait_time = between(5, 9)

from locust import HttpLocust, TaskSet, task, between

class MyTaskSet(TaskSet):
    @task(2)
    def queue(self):
        response = self.client.post("/v1/queue", {"user_id": 123,"cross_dispatch": 0,"order_id": "asd","initiated_at": "{{currentDate}}","status": "asdasd"})
        print("[/v1/queue] Response status code:", response.status_code)
        print("[/v1/queue] Response content:", response.text)

    @task(1)
    def callback(self):
        response = self.client.post("/callback", {"name":"Joe","email":"joe@labstack"})
        print("callback] Response status code:", response.status_code)
        print("callback] Response content:", response.text)

class MyLocust(HttpLocust):
    task_set = MyTaskSet
    wait_time = between(5, 15)