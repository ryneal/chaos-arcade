package com.github.ryneal.chaosarcade.config;

import io.kubernetes.client.openapi.ApiClient;
import io.kubernetes.client.openapi.apis.CoreV1Api;
import io.kubernetes.client.util.Config;
import lombok.extern.slf4j.Slf4j;
import org.apache.logging.log4j.util.Strings;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.io.IOException;
import java.util.Collections;
import java.util.List;
import java.util.Optional;

@Configuration
@Slf4j
public class KubernetesConfig {

    private final List<String> allowedNamespaces;

    public KubernetesConfig(@Value("${ALLOWED_NAMESPACES:}") List<String> allowedNamespaces) {
        this.allowedNamespaces = Optional.ofNullable(allowedNamespaces).orElseGet(Collections::emptyList);
        log.info("Allowed namespaces are: {}", Strings.join(this.allowedNamespaces, ','));
    }

    @Bean
    public ApiClient kubernetesClient() throws IOException {
        return Config.defaultClient();
    }

    @Bean
    public CoreV1Api kubernetesCoreV1Api(ApiClient kubernetesClient) {
        return new CoreV1Api(kubernetesClient);
    }

    @Bean
    public List<String> allowedNamespaces() {
        return this.allowedNamespaces;
    }

}
