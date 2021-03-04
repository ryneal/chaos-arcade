package com.github.ryneal.chaosarcade.service;

import com.github.ryneal.chaosarcade.model.Pod;
import io.kubernetes.client.openapi.ApiException;
import io.kubernetes.client.openapi.apis.CoreV1Api;
import io.kubernetes.client.openapi.models.V1Pod;
import io.kubernetes.client.openapi.models.V1PodList;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Optional;
import java.util.concurrent.ThreadLocalRandom;

@Service
@Slf4j
public class PodServiceImpl implements PodService {

    private final CoreV1Api coreV1Api;
    private final List<String> allowedNamespaces;

    public PodServiceImpl(CoreV1Api coreV1Api, List<String> allowedNamespaces) {
        this.coreV1Api = coreV1Api;
        this.allowedNamespaces = allowedNamespaces;
    }

    @Override
    public Optional<Pod> getRandomPod() {
        return this.getAllPods()
                .stream()
                .sorted((o1, o2) -> ThreadLocalRandom.current().nextInt(-1, 2))
                .findAny();
    }

    @Override
    public List<Pod> getAllPods() {
        try {
            List<Pod> results = new ArrayList<>();
            for (String ns : this.allowedNamespaces) {
                V1PodList v1PodList = this.coreV1Api.listNamespacedPod(ns, null, null,
                        null, null, null, null, null,
                        null, null, null);
                for (V1Pod v1Pod : v1PodList.getItems()) {
                    Optional.ofNullable(v1Pod)
                            .map(V1Pod::getMetadata)
                            .map(m -> Pod.builder()
                                    .name(m.getName())
                                    .namespace(m.getNamespace())
                                    .build())
                            .ifPresent(results::add);
                }
            }
            return results;
        } catch (ApiException e) {
            log.debug("failed to get all pods", e);
            return Collections.emptyList();
        }
    }

    @Override
    public Optional<Pod> deletePod(Pod pod) {
        try {
            return Optional.ofNullable(this.coreV1Api.deleteNamespacedPod(pod.getName(), pod.getNamespace(),
                    null, null, null,
                    null, null, null))
                    .map(V1Pod::getMetadata)
                    .map(m -> Pod.builder()
                            .namespace(m.getNamespace())
                            .name(m.getName())
                            .build());
        } catch (ApiException e) {
            log.debug("failed to delete pod", e);
            return Optional.empty();
        }
    }
}
