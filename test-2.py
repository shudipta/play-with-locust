import sys
from locust import Locust, HttpLocust, TaskSet, task, between, TaskSequence
from locust import events
from datetime import datetime, timezone
# from pytzdata import pytz
import pendulum
import time
import json

globala = -100


def on_my_event(locust_instance, exception, tb):
    # print("Event was fired with arguments: %s, %s" % (kw.get("locust_instance"),
    #                                                                            kw.get("exception")))
    tb.print_stack()
    # for arg in kw:
    #     if kw.get(arg):
    #         print("%s=%s" % (arg, kw.get(arg)))


class MyTaskSet(TaskSequence):
    my_event = events.locust_error

    def on_start(self):
        # self.locust.a = 0
        print("[on_start]: %d" % self.locust.a)
        # self.my_event += on_my_event
        # events.locust_error += on_my_event

    def on_stop(self):
        self.locust.a = 0
        print("[on_stop]: %d" % self.locust.a)

    @task(1)
    def my_task1(self):
        t1 = time.time()
        print("[my_task1]: %d" % self.locust.a)
        self.locust.a += 1

        testValue = None
        try: assert (testValue is not None)
        except:
        order = {
            "order_id": "order_id",
            "initiated_at": str(datetime.utcnow().isoformat("T")+"Z"),
            "status": "pre-accept",
            "created_at": str(datetime.utcnow().isoformat("T")+"Z")
        }
        print (json.dumps(order, indent=2))
        with testValue.post(
            "http://localhost:4200/v1/order/accepted",
            json=order,
            catch_response=True
        ) as ac_resp:
            if ac_resp.status_code != 201:
                ac_resp.failure("[/order/accepted] %s status_code: Expect 201, Got %d" % ("order_id", ac_resp.status_code))
            else:
                print("/order/accepted] %s: success" % "order_id")
                ac_resp.success()

        atest = 3
        #
        # self.my_event.fire(
        #             a=atest,
        #             b=5,
        #             # locust_instance=self.locust,
        #             # exception="Expect: 5, Got: %d" % 5,
        #         )
        # self.my_event.fire(
        #             a=atest,
        #             b=5,
        #             # locust_instance=self.locust,
        #             # exception="Expect: 5, Got: %d" % 5,
        #         )
        try:
            # 1 / 0
            t2 = time.time()
            assert (atest == 5)
        except:
            # tb = e.__traceback__  # or sys.exc_info()[2]
            # events.locust_error.fire(
            #     # a=atest,
            #     # b=5,
            #     locust_instance=self.locust,
            #     exception="Expect: 5, Got: %d" % atest,
            #     # tb=sys.exc_info()[2]
            #     tb=tb
            # )
            events.request_failure.fire(
                request_type="task",
                name="my_task",
                response_time=(t2-t1)*100,
                response_length=0,
                exception="Expect: 5, Got: %d" % atest
            )
        else:
            pass

        # if :
        #     e = Exception("Expect: 5, Got: %d" % 5)
        #     events.locust_error.fire(
        #         # a=atest,
        #         # b=5,
        #         locust_instance=self.locust,
        #         exception="Expect: 5, Got: %d" % 5,
        #         tb=e.__traceback__
        #         # tb=Traceback(sys.exc_info()[2])
        #     )
            # raise Exception("Expect: 5, Got: %d" % atest)
        # if atest == 5:
        #     self.my_event.fire(
        #         a=atest,
        #         b=5,
        #         # locust_instance=self.locust,
        #         # exception="Expect: 5, Got: %d" % 5,
        #     )



    # @task(2)
    # def my_task2(self):
    #     print("[my_task2]: %d" % (self.locust.a))
    #     self.locust.a += 1

class MyLocust(HttpLocust):
    a = globala
    print(datetime.utcnow())
    print(datetime(2020, 4, 9, 4, 11, 6))
    p = {"t": str(datetime(2020, 4, 9, 4, 11, 6))}
    print(json.dumps(p, indent=2))

    # @setup
    def setup(self):
        global globala

        globala = -5
        a = -5

    task_set = MyTaskSet
    wait_time = between(15, 15)

    print("stop: ", a)
