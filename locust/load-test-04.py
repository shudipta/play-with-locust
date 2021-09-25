import os
import sys

sys.path.append(os.getcwd())

import common.utils as myUtils
from locust import HttpLocust, TaskSet, task, between, constant
import pendulum
import random, string
import time

order_cnt = 0
base_order_name = ''.join(random.choice(string.ascii_lowercase) for i in range(5))
order_lifetime = int(os.environ.get('DISPATCHER_ORDER_LIFETIME', 60))
pair_limit = int(os.environ.get('DISPATCHER_PAIR_LIMIT', 5)) + 1
pair_checking_ignore_after = int(os.environ.get('DISPATCHER_PAIR_CHECKING_IGNORE_AFTER', 40))

mysql_pass = os.environ.get('DISPATCHER_MYSQL_ROOT_PASSWORD', '12345')
mysql_db = os.environ.get('DISPATCHER_MYSQL_DATABASE', 'dispatcher')
last_order_state_sql = """
SELECT statuses.status, statuses.created_at, details.driver_id
FROM (
    SELECT s.order_id, s.status, s.created_at
    FROM order_statuses s
    WHERE s.order_id = '{oid}'
    ORDER BY s.id DESC LIMIT 1
) statuses
INNER JOIN order_details details
ON statuses.order_id = details.order_id;
"""


class MyTaskSet(TaskSet):
    @task(1)
    def driverSkip(self):
        global order_cnt

        order_cnt += 1
        cur_order_offset = order_cnt
        # if cur_order_offset % pair_limit == 0:
        #     time.sleep(order_lifetime)

        user_id = cur_order_offset
        order_id = base_order_name + "-" + str(cur_order_offset)
        initiated_at = pendulum.now(tz="UTC").to_iso8601_string()
        payload = {
            "user_id": user_id,  # "user_id": 123
            "cross_dispatch": 0,
            "order_id": order_id,  # "order_id": "asdsd-123"
            "initiated_at": initiated_at,  # "initiated_at": "2020-02-24T07:43:04Z"
            "status": "asdasd"
        }
        myUtils.check_post_request(self.client, "http://localhost:4200", "/v1/queue", payload)
        t1 = time.time()

        print(">>>>>>>>> hey I am waiting for status changes of order_id " + str(order_id))
        status = ""
        it = 0
        skipped = False
        accepted = False
        sleep_time = 1
        while it < order_lifetime + 10:
            time.sleep(sleep_time)
            it += sleep_time

            res = myUtils.runQuery(
                pwd=mysql_pass, db=mysql_db, sql=last_order_state_sql.format(oid=order_id))
            if len(res) == 0:
                continue

            status = res[0][0]
            created_at = res[0][1]
            driver_id = res[0][2]
            print(">>>>>>>>>> (order, it, status) = (%s, %d, %s)" % (order_id, it, status))

            if status == "pre-accept":
                order = {
                    "user_id": user_id,
                    "order_id": order_id,
                    "initiated_at": initiated_at,
                    "status": status,
                    "created_at": str(created_at.isoformat("T") + "Z"),
                    "driver_id": int(driver_id),
                }
                if skipped is False:
                    skipped = myUtils.check_post_request(
                        self.client, "http://localhost:4200", "/v1/order/skipped", order)
                elif accepted is False:
                    accepted = myUtils.check_post_request(
                        self.client, "http://localhost:4200", "/v1/order/accepted", order)
                continue

            if status == "accepted":
                break

        resp_time = (time.time() - t1) * 1000
        print("+++++++++++++ (order, it, status) = (%s, %d, %s)" % (order_id, it, status))
        expcected_status = "accepted"

        try:
            assert (status == expcected_status)
        except AssertionError:
            myUtils.locust_event(
                request_type="TASK", name="driverSkip", response_time=resp_time, response_length=0,
                exception="%s: Order %s -> Expect '%s', Got '%s'" % (AssertionError, order_id, expcected_status, status)
            )
        else:
            myUtils.locust_event(
                request_type="TASK", name="driverSkip", response_time=resp_time, response_length=0)


class MyLocust(HttpLocust):
    task_set = MyTaskSet
    # wait_time = between(10, 15)
    wait_time = constant(500)
