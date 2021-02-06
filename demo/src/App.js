import React, { useState } from "react";
import "./App.css";
import "antd/dist/antd.css";
import { Form, Input, Button, Row, Col } from "antd";
import axios from "axios";
import { Alert } from "antd";
import { Image } from "antd";
import isURL from "validator/lib/isURL";

const App = () => {
  const [URL, setURL] = useState("https://www.gopl.io/");
  const [fetching, setFetching] = useState(false);
  const [images, setImages] = useState(undefined);
  const [errors, setErrors] = useState(undefined);

  const fetchImages = async (e, url) => {
    e.preventDefault();

    if (!isURL(url.trim())) {
      setErrors("Invalid email");
      setFetching(false);
      return;
    }
    if (!url.includes("http") && !url.includes("https")) {
      setErrors("Provide http or https");
      setFetching(false);
      return;
    }

    const config = {
      method: "post",
      url: "/images",
      headers: {
        "Content-Type": "application/json",
      },
      data: JSON.stringify({ url }),
    };

    try {
      const res = await axios(config);
      if (res.data) {
        setImages(res.data);
      }
    } catch (error) {
      if (error.response.data.message) {
        setErrors(error.response.data.message);
      } else {
        setErrors(error.response.data);
      }
    }
    setFetching(false);
  };

  return (
    <>
      <Row justify="center">
        <Form
          name="basic"
          size="large"
          defaultValue="https://www.gopl.io/"
          onValuesChange={(ch) => {
            setURL(ch.url);
          }}
        >
          <Form.Item
            label="URL"
            name="url"
            rules={[
              {
                required: true,
                message: "Please input your username!",
              },
            ]}
            initialValue="https://www.gopl.io/"
          >
            <Input />
          </Form.Item>

          <Form.Item>
            <Button
              type="primary"
              htmlType="submit"
              onClick={(e) => {
                setFetching(true);
                fetchImages(e, URL);
              }}
              style={{ width: "100%" }}
              loading={fetching}
            >
              Submit
            </Button>
          </Form.Item>
        </Form>
      </Row>
      {errors && <Alert message={errors} type="error" />}

      <Row justify="space-around">
        {images &&
          images.map((image, i) => (
            <Col>
              <Image preview={false} key={i} width={200} src={image.image} />
            </Col>
          ))}
      </Row>
    </>
  );
};

export default App;
