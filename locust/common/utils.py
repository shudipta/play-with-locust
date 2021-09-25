import json
import mysql.connector
from locust import events


def getUserID(pair_limit, order_cnt):
    """Calculates a number that will be treated as the userID

    Args:
    * order_cnt: type int, current order count
    
    :rtype: int
    """
    if order_cnt % pair_limit != 0:
        return int(order_cnt / pair_limit) + 1
    return int(order_cnt / pair_limit)


def runQuery(pwd="", db="", sql=None):
    dispatcher_db = mysql.connector.connect(
        host="localhost",
        port="33066",
        user="root",
        passwd=pwd,
        database=db,
    )
    mysql_cursor = dispatcher_db.cursor()
    mysql_cursor.execute(sql)
    return mysql_cursor.fetchall()


def check_post_request(client=None, host="", path="", payload=None):
    print(json.dumps(payload, indent=2))
    order_id = payload.get("order_id")

    if client is None:
        return False

    with client.post(
            host + path,  # "http://localhost:4200/v1/order/canceled",
            json=payload,
            catch_response=True
    ) as resp:
        if resp.status_code != 201:
            resp.failure("[%s] %s status_code: Expect 201, Got %d" % (path, order_id, resp.status_code))
            return False
        else:
            print("%s] %s: success" % (path, order_id))
            resp.success()
            return True


def check_get_request(client=None, host="", path="", debug_info={"order_id": "<not_provided>"}):
    order_id = debug_info.get("order_id")
    if client is None:
        return False

    with client.get(
            host + path,  # "http://localhost:4200/pair?user_id={uid}&driver_id={did}",
            catch_response=True
    ) as resp:
        if resp.status_code != 200:
            resp.failure("[%s] %s status_code: Expect 201, Got %d" % (resp, order_id, resp.status_code))
            return False
        else:
            print("%s] %s: success" % (path, order_id))
            resp.success()
            return True


def locust_event(
        request_type="TASK",
        name="",
        response_time=0,
        response_length=0,
        exception=None):
    if exception is None:
        events.request_success.fire(
            request_type=request_type,
            name=name,
            response_time=response_time,
            response_length=response_length
        )
    else:
        events.request_failure.fire(
            request_type=request_type,
            name=name,
            response_time=response_time,
            response_length=response_length,
            exception=exception
        )
