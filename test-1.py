from locust import HttpLocust, TaskSet, task, between, constant
# from datetime import datetime, timezone
import pendulum

order_cnt = 0

class MyTaskSet(TaskSet):
    # def on_start(self):
    #     """ on_start is called when a Locust start before any task is scheduled """
    #     print("[on_start]", self.locust)
    #
    # def on_stop(self):
    #     """ on_stop is called when the TaskSet is stopping """
    #     print("[on_stop]", self.locust)

    @task(2)
    def queue(self):
        global order_cnt
        order_cnt = order_cnt + 1
        user_id = order_cnt
        order_id = "asd" + str(order_cnt)
        # initiated_at = str(datetime.now(timezone.utc))
        initiated_at = pendulum.now().to_iso8601_string()
        # payload = {"user_id": 123, "cross_dispatch": 0, "order_id": "asd", "initiated_at": "2020-02-24T07:43:04+06:00", "status": "asdasd"}
        payload = {
            "user_id": user_id,
            "cross_dispatch": 0,
            "order_id": order_id,
            "initiated_at": initiated_at,
            "status": "asdasd"
        }
        print(">>>>>>>>>> payload: %s" % payload)
        with self.client.post(
            "http://localhost:4200/v1/queue",
            # headers,
            json=payload,
            catch_response=True
        ) as response:
            print("[/v1/queue] Response status code: '%d'" % response.status_code)
            print("[/v1/queue] Response content: '%s'" % response.text)
            print("[/v1/queue] Response content1: '%s'" % response.content)
            if response.status_code != 201:
                response.failure("[/v1/queue] status_code: Expect 201, Got " + str(response.status_code))
            else:
                print("[/v1/queue] hurray! success")
                response.success()

#     @task(1)
#     def callback(self):
#         payload = {"driver_id": 0, "user_id": 123, "order_id": "asd", "location": {"lat": 23.90, "lon": 90.23}}
#
#         # {"driver_id": 723, "user_id": 256, "order_id": "e27wbw", "location": {"lat": 23.90, "lon": 90.23}},
#         #####################################################
#         with self.client.post(
#             "http://localhost:4200/v1/callback",
#             # headers,
#             json=payload,
#             catch_response=True
#         ) as response:
#             print("[/callback] Response status code: '%d'" % response.status_code)
#             print("[/callback] Response test: '%s'" % response.text)
#             print("[/callback] Response content: '%s'" % response.content)
#             if response.text != "":
#                 response.failure("[/callback] Expect '', Got " + response.text)
#             else:
#                 print("[/callback] hurray! success")
#                 response.success()

class MyLocust(HttpLocust):
    task_set = MyTaskSet
    # wait_time = between(10, 15)
    wait_time = constant(500)
