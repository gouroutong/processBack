package com.xprocess.flowable;


import org.flowable.bpmn.BpmnAutoLayout;
import org.flowable.bpmn.converter.BpmnXMLConverter;
import org.flowable.bpmn.model.BpmnModel;
import org.flowable.common.engine.api.io.InputStreamProvider;
import org.flowable.common.engine.impl.util.io.InputStreamSource;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.io.ByteArrayInputStream;
import java.io.InputStream;

/**
 * Created by simod on 2020/3/11
 */
@RestController
@RequestMapping(value = "/api")
public class Api {
    public Api() {
    }

    @PostMapping("/calc_coord")
    public String generateCoord(@RequestBody String str) throws Exception {
        InputStream inputStream = new ByteArrayInputStream(str.getBytes());
        BpmnXMLConverter bpmnXMLConverter = new BpmnXMLConverter();
        InputStreamProvider inputStreamProvider = new InputStreamSource(inputStream);
        BpmnModel bpmnModel = bpmnXMLConverter.convertToBpmnModel(inputStreamProvider, true, false, "UTF-8");
        new BpmnAutoLayout(bpmnModel).execute();
        byte[] bytes = bpmnXMLConverter.convertToXML(bpmnModel);
        String xmlContenxt = new String(bytes, "utf-8");
        return xmlContenxt;
    }
}

